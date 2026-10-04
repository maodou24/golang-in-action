package factory_method

import "testing"

func TestFactoryMethodPlus(t *testing.T) {
	if Calculate(&PlusFactory{}, 1, 2) != 3 {
		t.Fail()
	}
}

func TestFactoryMethodSub(t *testing.T) {
	if Calculate(&SubFactory{}, 4, 2) != 2 {
		t.Fail()
	}
}
