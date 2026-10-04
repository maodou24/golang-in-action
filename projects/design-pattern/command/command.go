package command

import "fmt"

type Command interface {
	Execute()
}

type CommandA struct {
	receiver Receiver
}

func NewCommandA(receiver Receiver) *CommandA {
	return &CommandA{receiver}
}

func (c *CommandA) Execute() {
	fmt.Println("Execute Command A, print something")
	c.receiver.Print()
}

type Receiver struct {
	name string
}

func (r *Receiver) Print() {
	fmt.Printf("%v do some work\n", r.name)
}

type Invoker struct {
	command []Command
}

func (i *Invoker) Add(command Command) {
	i.command = append(i.command, command)
}

func (i *Invoker) invoker() {
	for _, command := range i.command {
		command.Execute()
	}
}
