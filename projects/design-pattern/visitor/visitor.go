package visitor

import "fmt"

type Visitor interface {
	Visit(Element)
}

type ConcreteVisitor struct {
}

func (c *ConcreteVisitor) Visit(e Element) {
	switch e.(type) {
	case *EleA:
		fmt.Println("A")
	case *EleB:
		fmt.Println("B")
	}
}

type Element interface {
	Accept(Visitor)
}

type ElementCol struct {
	elements []Element
}

func (c *ElementCol) Add(ele Element) {
	c.elements = append(c.elements, ele)
}

func (c *ElementCol) Accept(visitor Visitor) {
	for _, ele := range c.elements {
		ele.Accept(visitor)
	}
}

type EleA struct{}

func (e *EleA) Accept(v Visitor) {
	v.Visit(e)
}

type EleB struct{}

func (e *EleB) Accept(v Visitor) {
	v.Visit(e)
}
