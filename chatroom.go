package main

type ChatRoom struct {
	chatters map[*Chatter]bool

	joinChatRoom chan *Chatter

	leaveChatRoom chan *Chatter

	forwardMessages chan []byte
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		chatters:        make(map[*Chatter]bool),
		joinChatRoom:    make(chan *Chatter),
		leaveChatRoom:   make(chan *Chatter),
		forwardMessages: make(chan []byte),
	}
}

func (cr *ChatRoom) StartChat() {
	for {
		select {
		case chatter := <-cr.joinChatRoom:
			cr.chatters[chatter] = true
		case chatter := <-cr.leaveChatRoom:
			delete(cr.chatters, chatter)
			close(chatter.receive)
		}
	}
}
