package strategy

import "fmt"

type Strategy interface {
	DoOperation()
}

type StrategyA struct{}

func (s *StrategyA) DoOperation() {
	fmt.Println("A")
}

type StrategyB struct{}

func (s *StrategyB) DoOperation() {
	fmt.Println("B")
}

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) Execute() {
	c.strategy.DoOperation()
}
