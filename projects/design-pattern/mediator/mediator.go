package mediator

import (
	"fmt"
	"time"
)

type User struct {
	name string
}

func (u User) SendMsg(msg string) {
	GetChatRoom().showMessage(u, msg)
}

type ChatRoom struct{}

var chatRoom = ChatRoom{}

func GetChatRoom() ChatRoom {
	return chatRoom
}

func (c ChatRoom) showMessage(user User, msg string) {
	fmt.Printf("%v %s: %s\n", time.Now(), user.name, msg)
}
