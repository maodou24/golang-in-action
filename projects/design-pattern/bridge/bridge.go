package bridge

import "fmt"

type Car interface {
	Driver()
}

type JeepCar struct {
	Engine
}

func NewJeepCar(e Engine) JeepCar {
	return JeepCar{Engine: e}
}

func (f JeepCar) Driver() {
	fmt.Println("Driver Jeep Car")
	f.Engine.Start()
}

type Engine interface {
	Start()
}

type Fuel struct{}

func (f Fuel) Start() {
	fmt.Println("Fuel Engine start")
}

type Electric struct{}

func (f Electric) Start() {
	fmt.Println("Electric Engine start")
}
