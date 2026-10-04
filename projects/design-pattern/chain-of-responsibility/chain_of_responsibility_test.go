package chain_of_responsibility

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	client := &Client{}

	handler1 := &FirstHandler{}
	handler2 := &SecondHandler{}

	handler1.SetNext(handler2)

	client.SetHandler(handler1)
	client.ReceiverReq("url1")
}
