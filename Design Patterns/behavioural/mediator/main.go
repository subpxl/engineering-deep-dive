package main

import "fmt"

func main() {
	fmt.Println("mediator pattern")

	stationManager := NewStationManager()

	passengerTrain := NewPassengerTrain(stationManager)
	freightTrain := NewFreightTrain(stationManager)

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()

}
