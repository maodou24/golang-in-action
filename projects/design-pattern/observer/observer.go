package observer

import "fmt"

type Subject interface {
	RegisterObserver(observer Observer)
}

type ConcreteSubject struct {
	context   string
	observers []Observer
}

func (c *ConcreteSubject) RegisterObserver(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *ConcreteSubject) UpdateContext(context string) {
	c.context = context
	c.notify()
}

func (c *ConcreteSubject) notify() {
	for i := range c.observers {
		c.observers[i].Update(c.context)
	}
}

type Observer interface {
	Update(string)
}

type ConcreteObserver struct {
	name string
}

func (c ConcreteObserver) Update(ctx string) {
	fmt.Println(c.name, ctx)
}
