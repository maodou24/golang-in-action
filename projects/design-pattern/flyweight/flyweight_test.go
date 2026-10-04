package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	factory := NewFlyweightFactory()

	a := factory.Get("A")
	a.Set("A")
	if a.Do() != "A" {
		t.Error("A")
	}
}

func TestFlyweightTheSame(t *testing.T) {
	factory := NewFlyweightFactory()

	a := factory.Get("A")
	newA := factory.Get("A")
	if a != newA {
		t.Error("A not the same")
	}
}
