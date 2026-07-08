package main

import "fmt"

type PassengerTrain struct {
	mediator Mediator
}

func NewPassengerTrain(mediator Mediator) *PassengerTrain {
	return &PassengerTrain{
		mediator: mediator,
	}
}

func (p *PassengerTrain) Arrive() {
	if p.mediator.canArrive(p) {
		fmt.Println("Passenger train has arrived at the station.")
		return
	}
	fmt.Println("Passenger train cannot arrive at the station. Waiting for clearance.")
}

func (p *PassengerTrain) Depart() {
	fmt.Println("Passenger train is departing from the station.")
	p.mediator.notifyAboutDeparture()
}

func (p *PassengerTrain) permitArrival() {
	fmt.Println("Passenger train is permitted to arrive at the station.")
	p.Arrive()
}
