package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user
type client struct {

	// a socket is the web socket for this user
	socket *websocket.Conn

	//receive is a channel to receive messages from other clients
	receive chan []byte

	//room is the room this client is chatting in
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
	defer c.socket.Close()
	// endlessly read messages from input
	for {
		_, msg, err := c.socket.ReadMessage()
		// break if there is an error
		if err != nil {
			return
		}

		outgoing := map[string]string{
			"name":    c.name,
			"message": string(msg),
		}

		jsMessage, err := json.Marshal(outgoing)
		if err != nil {
			fmt.Println("Enconding failed!")
			continue
		}

		// forward the message to the room
		c.room.forward <- jsMessage
	}
}

// ------- write()

// 1.Listens on c.receive channel.
// 2.Sends messages back to the WebSocket client.

func (c *client) write() {
	defer c.socket.Close()

	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)

		if err != nil {
			return
		}
	}
}
