package state

import "fmt"

type State interface {
	DoAction(ctx Context)
}

type Context interface {
	SetState(state State)
	GetState() State
}

func NewContext() Context {
	return &ContextImpl{state: nil}
}

type ContextImpl struct {
	state State
}

func (c *ContextImpl) SetState(state State) {
	c.state = state
}

func (c *ContextImpl) GetState() State {
	return c.state
}

type StateA struct{}

func (s *StateA) DoAction(ctx Context) {
	fmt.Println("in state A")
	ctx.SetState(s)
}

type StateB struct{}

func (s *StateB) DoAction(ctx Context) {
	fmt.Println("in state B")
	ctx.SetState(s)
}
