package main

type Train interface {
	Arrive()
	Depart()
	permitArrival()
}
