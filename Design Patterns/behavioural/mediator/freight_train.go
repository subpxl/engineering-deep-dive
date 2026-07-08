package main

import "fmt"

type FreightTrain struct {
	mediator Mediator
}

func NewFreightTrain(mediator Mediator) *FreightTrain {
	return &FreightTrain{
		mediator: mediator,
	}
}

func (f *FreightTrain) Arrive() {
	if f.mediator.canArrive(f) {
		fmt.Println("Freight train has arrived at the station.")
		return
	}
	fmt.Println("Freight train cannot arrive at the station. Waiting for clearance.")

}

func (f *FreightTrain) Depart() {
	fmt.Println("Freight train is departing from the station.")
	f.mediator.notifyAboutDeparture()
}

func (f *FreightTrain) permitArrival() {
	fmt.Println("Freight train is permitted to arrive at the station.")
	f.Arrive()
}
