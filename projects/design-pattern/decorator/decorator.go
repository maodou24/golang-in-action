package decorator

import "fmt"

type Component interface {
	Do() string
}

type ConcreteComponent struct{}

func (a ConcreteComponent) Do() string {
	return "first"
}

type Decorator struct {
	Component
}

func NewDecorator(component Component) Component {
	return &Decorator{component}
}

func (a Decorator) Do() string {
	r := a.Component.Do() + " second"
	fmt.Println(r)
	return r
}
