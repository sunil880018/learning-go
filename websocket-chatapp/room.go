package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type room struct {

	// holds all current clients in the room
	clients map[*client]bool

	// join is a channel for all clients wishing to join the room
	join chan *client

	// leave is a channel for all clients wishing to leave the room
	leave chan *client

	// forward is a channel that holds incoming messages that should be forwarded to the other clients.
	forward chan []byte

	name string
}

// Redis client
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // default Redis
})

var ctx = context.Background()

// Creates a new chat room with empty clients map and channels.
func newRoom(name string) *room {
	r := &room{
		name:    name,
		forward: make(chan []byte, 256), // buffered to reduce blocking
		join:    make(chan *client, 64),
		leave:   make(chan *client, 64),
		clients: make(map[*client]bool),
	}

	// Start listening to Redis pub/sub for this room
	go r.subscribeRedis()

	return r
}

// subscribeRedis listens to Redis channel and pushes incoming messages into room.forward
func (r *room) subscribeRedis() {
	pubsub := rdb.Subscribe(ctx, r.name)
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		select {
		case r.forward <- []byte(msg.Payload):
			// delivered to room loop
		default:
			// if forward channel full, drop to avoid blocking
			log.Printf("[Room:%s] Dropped Redis message (forward channel full)", r.name)
		}
	}
}

// ------run()

// Infinite loop that reacts to events:
// join: add client to room.
// leave: remove client from room, close their channel.
// forward: broadcast a message to all connected clients.
func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// safely add new client
			if _, exists := r.clients[client]; !exists {
				r.clients[client] = true
				log.Printf("[Room:%s] Client %s joined | total: %d",
					r.name, client.name, len(r.clients))
			}

		case client := <-r.leave:
			// safely remove client
			if _, exists := r.clients[client]; exists {
				delete(r.clients, client)
				close(client.receive)
				log.Printf("[Room:%s] Client %s left | total: %d",
					r.name, client.name, len(r.clients))
			}

		case msg := <-r.forward:
			// publish to Redis so other servers can see it
			if err := rdb.Publish(ctx, r.name, msg).Err(); err != nil {
				log.Printf("[Room:%s] Redis publish error: %v", r.name, err)
				continue
			}

			// broadcast locally too
			for client := range r.clients {
				select {
				case client.receive <- msg: // non-blocking send
				default:
					// client channel is full → drop message to avoid blocking room
					log.Printf("[Room:%s] Dropping message for %s (slow client)",
						r.name, client.name)
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

var rooms = make(map[string]*room)
var mu sync.Mutex

// getRoom returns an existing room or creates a new one
func getRoom(name string) *room {
	mu.Lock()
	defer mu.Unlock()
	if r, ok := rooms[name]; ok {
		return r
	}
	r := newRoom(name)
	rooms[name] = r
	go r.run()
	return r
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	roomName := req.URL.Query().Get("room")
	if roomName == "" {
		http.Error(w, "Room name required", http.StatusBadRequest)
		return
	}

	realRoom := getRoom(roomName)

	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	client := &client{
		socket:  socket,
		receive: make(chan []byte, messageBufferSize),
		room:    realRoom,
		name:    fmt.Sprintf("user_%d", rand.Intn(1000)),
	}

	realRoom.join <- client
	defer func() { realRoom.leave <- client }()
	go client.write()
	client.read()
}

// 1.Added rdb = redis.NewClient(...) for Redis connection.
// 2.newRoom now calls subscribeRedis to listen for messages from Redis.
// 3.When a client sends a message → it goes to r.forward → room publishes it to Redis.
// 4.When Redis delivers a message → it’s broadcast to all local clients in that room.
