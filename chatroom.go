package main

import (
	"github.com/gorilla/websocket"
)

const (
	websocketBufferSize = 1024
	messagesBufferSize  = 1024
)

var (
	upgrader = websocket.Upgrader{HandshakeTimeout: websocketBufferSize, ReadBufferSize: messagesBufferSize}
)

type ChatRoom struct {
	// chatters are people chatting in the ChatRoom.
	chatters map[*Chatter]bool

	// joinChatRoom is a channel that admits new chatters into the ChatRoom.
	joinChatRoom chan *Chatter

	// leaveChatRoom is a channel that removes existing chatters from the ChatRoom.
	leaveChatRoom chan *Chatter

	// forwardMessages is a channel that allows forwarding functionality to all chatters.
	forwardMessages chan []byte
}

// NewChatRoom creates a new chat room.
func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		chatters:        make(map[*Chatter]bool),
		joinChatRoom:    make(chan *Chatter),
		leaveChatRoom:   make(chan *Chatter),
		forwardMessages: make(chan []byte),
	}
}

// StartChat starts a chat session.
func (cr *ChatRoom) StartChat() {
	for {
		select {
		case chatter := <-cr.joinChatRoom:
			cr.chatters[chatter] = true
		case chatter := <-cr.leaveChatRoom:
			delete(cr.chatters, chatter) // remove the chatter
			close(chatter.receive)       // close the chatter's access to the chatroom.
		case forwardedMessages := <-cr.forwardMessages:
			for chatter := range cr.chatters {
				chatter.receive <- forwardedMessages
			}

		}
	}
}
