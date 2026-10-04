package mediator

import "testing"

func TestMediator(t *testing.T) {
	user1 := User{name: "user1"}
	user2 := User{name: "user2"}

	user1.SendMsg("How are you?")
	user2.SendMsg("I'm fine")
}
