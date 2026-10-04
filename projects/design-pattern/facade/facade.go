package facade

type CarBuilder struct {
	jeep   JeepCar
	supper SuperCar
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{
		jeep:   JeepCar{},
		supper: SuperCar{},
	}
}

func (c *CarBuilder) BuildJeepCar() string {
	return c.jeep.Build()
}

func (c *CarBuilder) BuildSuperCar() string {
	return c.supper.Build()
}

type JeepCar struct{}

func (v *JeepCar) Build() string {
	return "Jeep Car"
}

type SuperCar struct{}

func (s *SuperCar) Build() string {
	return "Super Car"
}
