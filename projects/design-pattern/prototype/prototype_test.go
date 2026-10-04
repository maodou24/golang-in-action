package prototype

import (
	"testing"
)

func TestPrototypeCloneable(t *testing.T) {
	manager := NewPrototypeManager()

	a := &A{Name: "A"}
	manager.Set("A", a)
	newA := manager.Get("A").Clone()
	if a == newA {
		t.Errorf("A was not cloned properly")
	}
}

type A struct {
	Name string
}

func (a *A) Clone() Cloneable {
	r := *a
	return &r
}
