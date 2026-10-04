package simple_factory

import "testing"

func TestSimpleFactoryCar(t *testing.T) {
	jeepCar := NewCar("Jeep")
	if jeepCar.Run() != "Jeep Car" {
		t.Fatal()
	}

	superCar := NewCar("Super")
	if superCar.Run() != "Super Car" {
		t.Fatal()
	}
}
