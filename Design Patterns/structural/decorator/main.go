package main

import "fmt"

func main() {

	coffee := NewSimpleCoffee()
	lattee := NewLatte(coffee)

	sugarCoffee := NewSugarDecorator(coffee)

	fmt.Println(coffee.GetDescription(),coffee.GetPrice())
	
	fmt.Println()
	fmt.Println(lattee.GetDescription(),lattee.GetPrice())
	fmt.Println()

	fmt.Println(sugarCoffee.GetDescription(), sugarCoffee.GetPrice())

}

type Coffee interface {
	GetDescription() string
	GetPrice() float64
}

type SimpleCoffee struct {
	price float64
}

func NewSimpleCoffee() *SimpleCoffee {
	return &SimpleCoffee{price: 10.0}
}

func (c *SimpleCoffee) GetDescription() string {
	return "simple coffee"
}

func (c *SimpleCoffee) GetPrice() float64 {
	return c.price
}

type LatteeDecorator struct {
	inner Coffee
}

func NewLatte(inner Coffee) *LatteeDecorator {
	return &LatteeDecorator{inner: inner}
}

func (c *LatteeDecorator) GetPrice() float64 {
	return c.inner.GetPrice() + 5.0
}

func (c *LatteeDecorator) GetDescription() string {
	return c.inner.GetDescription() + "  lattee"
}

type SugarDecorator struct {
	inner Coffee
}

func NewSugarDecorator(inner Coffee) *SugarDecorator {
	return &SugarDecorator{
		inner: inner,
	}
}

func (c *SugarDecorator) GetPrice() float64      { return c.inner.GetPrice() + 10.0 }
func (c *SugarDecorator) GetDescription() string { return c.inner.GetDescription() + " added sugar" }
