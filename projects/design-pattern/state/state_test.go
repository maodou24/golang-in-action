package state

import "testing"

func TestState(t *testing.T) {
	ctx := NewContext()

	a := &StateA{}
	a.DoAction(ctx)

	b := &StateB{}
	b.DoAction(ctx)
}
