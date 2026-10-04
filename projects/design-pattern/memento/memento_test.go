package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	caretaker := &Caretaker{}

	originator := &Originator{}

	originator.SetState("A")
	caretaker.Add(originator.SaveToMemento())
	fmt.Println(originator.GetState())

	originator.SetState("B")
	caretaker.Add(originator.SaveToMemento())
	fmt.Println(originator.GetState())

	originator.RecoverFromMemento(caretaker.Get(0))
	fmt.Println(originator.GetState())

	originator.RecoverFromMemento(caretaker.Get(1))
	fmt.Println(originator.GetState())
}
