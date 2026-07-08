package main

import "fmt"

func main() {

	vm := NewVendingMachine()
	vm.SelectItem("A1")
	vm.InsertCoin(1.0)
	vm.DispenseItem()
}

type MachineState interface {
	SelectItem(context *VendingMachine, itemCode string)
	InsertCoin(context *VendingMachine, amount float64)
	DispenseItem(context *VendingMachine)
}

type IdleState struct{}

func (s *IdleState) SelectItem(context *VendingMachine, itemCode string) {
	fmt.Println("Item selected:", itemCode)
	context.SetSelectedItem(itemCode)
	context.SetState(&ItemSelectedState{})
}

func (s *IdleState) InsertCoin(context *VendingMachine, amount float64) {
	fmt.Println("Please select an item first before inserting coin")
}

func (s *IdleState) DispenseItem(context *VendingMachine) {
	fmt.Println("no item selected, nothing to dispense")
}

type ItemSelectedState struct{}

func (s *ItemSelectedState) SelectItem(context *VendingMachine, itemCode string) {
	fmt.Println("Item already selected, please insert coin")
}

func (s *ItemSelectedState) InsertCoin(context *VendingMachine, amount float64) {
	fmt.Println("Coin inserted:", amount)
	context.SetInsertedAmount(amount)
	context.SetState(&HasMoneyState{})
}

func (s *ItemSelectedState) DispenseItem(context *VendingMachine) {
	fmt.Println("Please insert coin first before dispensing item")
}

type HasMoneyState struct{}

func (s *HasMoneyState) SelectItem(context *VendingMachine, itemCode string) {
	fmt.Println("Item already selected, please wait for dispensing")
}

func (s *HasMoneyState) InsertCoin(context *VendingMachine, amount float64) {
	fmt.Println("Coin already inserted, please wait for dispensing")
}

func (s *HasMoneyState) DispenseItem(context *VendingMachine) {
	fmt.Println("Dispensing item:", context.GetSelectedItem())
	context.SetState(&DispensingState{})
	fmt.Println("item dispensed successfully")
	context.Reset()
}

type DispensingState struct{}

func (s *DispensingState) SelectItem(context *VendingMachine, itemCode string) {
	fmt.Println("Dispensing in progress, please wait")
}

func (s *DispensingState) InsertCoin(context *VendingMachine, amount float64) {
	fmt.Println("Dispensing in progress, please wait")
}

func (s *DispensingState) DispenseItem(context *VendingMachine) {
	fmt.Println("Dispensing in progress, please wait")
}

// implement context

type VendingMachine struct {
	currentState   MachineState
	selectedItem   string
	insertedAmount float64
}

func NewVendingMachine() *VendingMachine {
	return &VendingMachine{
		currentState: &IdleState{},
	}
}

func (v *VendingMachine) SetState(newState MachineState) {
	v.currentState = newState
}

func (v *VendingMachine) SetSelectedItem(itemCode string) {
	v.selectedItem = itemCode
}

func (v *VendingMachine) SetInsertedAmount(amount float64) {
	v.insertedAmount = amount
}

func (v *VendingMachine) GetSelectedItem() string {
	return v.selectedItem
}

func (v *VendingMachine) InsertCoin(amount float64) {
	v.currentState.InsertCoin(v, amount)
}

func (v *VendingMachine) DispenseItem() {
	v.currentState.DispenseItem(v)
}

func (v *VendingMachine) SelectItem(itemCode string) {
	v.currentState.SelectItem(v, itemCode)
}

func (v *VendingMachine) Reset() {
	v.selectedItem = ""
	v.insertedAmount = 0
	v.currentState = &IdleState{}
}
