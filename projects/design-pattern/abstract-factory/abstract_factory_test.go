package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	factoryA := NewCarFactory("A")
	factoryA.Run()
	factoryA.Light()

	factoryB := NewCarFactory("B")
	factoryB.Run()
	factoryB.Light()
}
