package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	meal := NewMeal()
	meal.Add(&Coke{})
	meal.Add(&Pepsi{})
	meal.Add(&VegBurger{})

	meal.ShowItems()
	fmt.Println(meal.GetPrice())
}
