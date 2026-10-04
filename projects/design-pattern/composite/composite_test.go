package composite

import "testing"

func TestComponent(t *testing.T) {
	c := NewComponent("c")

	c1 := NewComponent("c1")
	c2 := NewComponent("c2")
	c.Add(c1)
	c.Add(c2)

	c.Print()
}

func TestComponentLeaf(t *testing.T) {
	var c Component = &Leaf{}

	c.SetName("leaf")

	c.Add(NewComponent("c1"))
	c.Add(NewComponent("c2"))
	c.Print()
}
