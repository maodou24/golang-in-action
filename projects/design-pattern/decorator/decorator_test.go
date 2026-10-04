package decorator

import "testing"

func TestComponent(t *testing.T) {
	d := NewDecorator(ConcreteComponent{})
	if d.Do() != "first second" {
		t.Error("expecting 'first second'")
	}
}
