package visitor

import "testing"

func TestVisitor(t *testing.T) {
	col := &ElementCol{}

	col.Add(&EleA{})
	col.Add(&EleB{})

	visitor := &ConcreteVisitor{}
	col.Accept(visitor)
}
