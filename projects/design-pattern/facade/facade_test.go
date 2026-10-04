package facade

import (
	"testing"
)

func TestFacadeJeep(t *testing.T) {
	carBuilder := NewCarBuilder()

	if carBuilder.BuildJeepCar() != "Jeep Car" {
		t.Fail()
	}
}

func TestFacadeSuper(t *testing.T) {
	carBuilder := NewCarBuilder()

	if carBuilder.BuildSuperCar() != "Super Car" {
		t.Fail()
	}
}
