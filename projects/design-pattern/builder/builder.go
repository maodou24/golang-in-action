package builder

import "fmt"

type Builder interface {
	Name() string
	Packing() string
	Price() float32
}

type Meal struct {
	items []Builder
}

func NewMeal() *Meal {
	return &Meal{}
}

func (m *Meal) Add(builder Builder) {
	m.items = append(m.items, builder)
}

func (m *Meal) GetPrice() float32 {
	var result float32
	for _, builder := range m.items {
		result += builder.Price()
	}
	return result
}

func (m *Meal) ShowItems() {
	for _, builder := range m.items {
		fmt.Printf("item: %v, packing: %v, price: %v\n",
			builder.Name(), builder.Packing(), builder.Price())
	}
}

type VegBurger struct {
}

func (v *VegBurger) Name() string {
	return "VegBurger"
}

func (v *VegBurger) Packing() string {
	return "Wrapper"
}

func (v *VegBurger) Price() float32 {
	return 1.25
}

type ChickenBurger struct {
}

func (v *ChickenBurger) Name() string {
	return "ChickenBurger"
}

func (v *ChickenBurger) Packing() string {
	return "Wrapper"
}

func (v *ChickenBurger) Price() float32 {
	return 2.75
}

type Coke struct {
}

func (v *Coke) Name() string {
	return "Coke"
}

func (v *Coke) Packing() string {
	return "Bottle"
}

func (v *Coke) Price() float32 {
	return 3
}

type Pepsi struct {
}

func (v *Pepsi) Name() string {
	return "Pepsi"
}

func (v *Pepsi) Packing() string {
	return "Bottle"
}

func (v *Pepsi) Price() float32 {
	return 3
}
