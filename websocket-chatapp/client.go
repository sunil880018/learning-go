package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user
type client struct {

	// a socket is the web socket for this user
	socket *websocket.Conn

	// receive is a channel to receive messages from other clients
	receive chan []byte

	// room is the room this client is chatting in
	room *room

	name string
}

// ----- read()

// 1.Listens for incoming messages from the client’s WebSocket.
// 2.Converts message into JSON with username + text.
// 3.Sends it to the room’s forward channel → so the room can broadcast it.

// Used to send messages from frontend to backend
func (c *client) read() {
	// close the connection when we are done
	defer func() {
		c.room.leave <- c // ensure cleanup
		c.socket.Close()
	}()

	// endlessly read messages from input
	for {
		_, msg, err := c.socket.ReadMessage()
		// break if there is an error (client disconnected)
		if err != nil {
			log.Printf("[Client:%s] Read error: %v", c.name, err)
			return
		}

		outgoing := map[string]string{
			"name":    c.name,
			"message": string(msg),
		}

		jsMessage, err := json.Marshal(outgoing)
		if err != nil {
			log.Printf("[Client:%s] Encoding failed: %v", c.name, err)
			continue
		}

		// forward the message to the room
		select {
		case c.room.forward <- jsMessage:
			// success
		default:
			// if forward channel is full, drop the message to avoid blocking
			log.Printf("[Client:%s] Dropping outgoing message (room forward channel full)", c.name)
		}
	}
}

// ------- write()

// 1.Listens on c.receive channel.
// 2.Sends messages back to the WebSocket client.
func (c *client) write() {
	defer func() {
		c.socket.Close()
	}()

	for msg := range c.receive {
		// write with error handling
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Printf("[Client:%s] Write error: %v", c.name, err)
			return
		}
	}
}
