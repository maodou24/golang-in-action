package abstract_factory

import "fmt"

type CarFactory interface {
	WheelFactory
	LampFactory
}

func NewCarFactory(carType string) CarFactory {
	if carType == "A" {
		return &ACarFactory{}
	} else if carType == "B" {
		return &BCarFactory{}
	}
	return nil
}

type ACarFactory struct {
	LampA
	WheelA
}

type BCarFactory struct {
	LampB
	WheelB
}

type WheelFactory interface {
	Run()
}

type LampFactory interface {
	Light()
}

type LampA struct {
}

func (LampA) Light() {
	fmt.Println("A Lamp")
}

type LampB struct {
}

func (LampB) Light() {
	fmt.Println("B Lamp")
}

type WheelA struct {
}

func (WheelA) Run() {
	fmt.Println("A Run")
}

type WheelB struct {
}

func (WheelB) Run() {
	fmt.Println("B Run")
}
