package main

import (
	"github.com/gorilla/websocket"
)

type Chatter struct {
	socket *websocket.Conn

	receive chan []byte

	chatroom *ChatRoom
}

func (c *Chatter) ReadMessages() {
	defer c.socket.Close()
	for {
		// continuous loop of reading
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.chatroom.forwardMessages <- message
	}
}

func (c *Chatter) WriteMessages() {
	defer c.socket.Close()
	for message := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}
