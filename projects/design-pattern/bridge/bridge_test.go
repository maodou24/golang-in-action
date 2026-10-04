package bridge

import "testing"

func TestBridgeElectric(t *testing.T) {
	jeepCar := NewJeepCar(Electric{})

	jeepCar.Driver()
}

func TestBridgeFuel(t *testing.T) {
	jeepCar := NewJeepCar(Fuel{})

	jeepCar.Driver()
}
