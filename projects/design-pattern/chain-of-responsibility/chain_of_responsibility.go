package chain_of_responsibility

import "fmt"

type Handler interface {
	SetNext(next Handler)
	Handle(string)
}

type FirstHandler struct {
	next Handler
}

func (c *FirstHandler) SetNext(next Handler) {
	c.next = next
}

func (c *FirstHandler) Handle(url string) {
	fmt.Println("First handle url:", url)
	c.next.Handle(url)
}

type SecondHandler struct {
	next Handler
}

func (c *SecondHandler) SetNext(next Handler) {
	c.next = next
}

func (c *SecondHandler) Handle(url string) {
	fmt.Println("Second handle url:", url)

	if c.next != nil {
		c.next.Handle(url)
	}
}

type Client struct {
	handler Handler
}

func (c *Client) SetHandler(h Handler) {
	c.handler = h
}

func (c *Client) ReceiverReq(url string) {
	c.handler.Handle(url)
}
