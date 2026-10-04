package composite

import "fmt"

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	Add(component Component)
	GetComponents() []Component
	Print()
}

func NewComponent(name string) Component {
	c := &Composite{}
	c.SetName(name)
	return c
}

type Leaf struct {
	parent Component
	name   string
}

func (l *Leaf) Parent() Component {
	return l.parent
}

func (l *Leaf) SetParent(c Component) {
	l.parent = c
}

func (l *Leaf) Name() string {
	return l.name
}

func (l *Leaf) SetName(name string) {
	l.name = name
}

func (l *Leaf) Add(component Component) {
}

func (l *Leaf) GetComponents() []Component {
	return nil
}

func (l *Leaf) Print() {
	fmt.Println(l.name)
}

type Composite struct {
	Leaf
	components []Component
}

func (c *Composite) Add(component Component) {
	c.components = append(c.components, component)
}

func (c *Composite) GetComponents() []Component {
	return c.components
}

func (c *Composite) Print() {
	fmt.Println(c.name)

	for _, component := range c.GetComponents() {
		component.Print()
	}
}
