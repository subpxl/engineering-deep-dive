package main

import "fmt"

type Power interface {
	SetName(name string)
	GetName() string
	SetPower(power int)
	GetPower() int
}

type Punch struct {
	name  string
	power int
}

func (p *Punch) SetName(name string) { p.name = name }
func (p *Punch) GetName() string     { return p.name }
func (p *Punch) SetPower(power int)  { p.power = power }
func (p *Punch) GetPower() int       { return p.power }

type TurboPunch struct {
	Punch
}

func NewTorboPunch() *TurboPunch {
	return &TurboPunch{
		Punch: Punch{name: "turbo punch", power: 100},
	}
}

type MegaPunch struct {
	Punch
}

func NewMegaPunch() *MegaPunch {
	return &MegaPunch{
		Punch: Punch{name: "mega punch", power: 200},
	}
}

func main() {

	turboPunch := NewTorboPunch()

	fmt.Println(turboPunch.GetName(), turboPunch.GetPower())

	megaPunch := NewMegaPunch()
	fmt.Println(megaPunch.GetName(), megaPunch.GetPower())

}
