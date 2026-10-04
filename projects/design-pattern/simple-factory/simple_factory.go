package simple_factory

func NewCar(name string) Car {
	switch name {
	case "Jeep":
		return &JeepCar{}
	case "Super":
		return &SuperCar{}
	default:
		return nil
	}
}

type Car interface {
	Run() string
}

type JeepCar struct{}

func (JeepCar) Run() string {
	return "Jeep Car"
}

type SuperCar struct{}

func (SuperCar) Run() string {
	return "Super Car"
}
