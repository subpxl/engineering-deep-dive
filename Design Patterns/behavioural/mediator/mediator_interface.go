package main

type Mediator interface {
	canArrive(train Train) bool
	notifyAboutDeparture()
}
