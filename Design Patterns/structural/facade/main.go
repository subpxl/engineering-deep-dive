package main

import "fmt"

func main() {

	amp := &Amplifier{}
	dvd := DVDPlayer{}
	projector := &Projector{}
	lights := &SmartLights{}
	streaming := &StreamingService{}

	theater := NewHomeTheaterFacade(amp, &dvd, projector, lights, streaming)
	theater.WatchMovie("titanic")
	fmt.Println("")

	theater.EndMovie()
}

type Amplifier struct{}

func (a *Amplifier) On()                  { fmt.Println("amplifier: powering on") }
func (a *Amplifier) Off()                 { fmt.Println("amplifier: powering off") }
func (a *Amplifier) SetVolume(volume int) { fmt.Println("amplifier: volume = ", volume) }

type DVDPlayer struct{}

func (d *DVDPlayer) On()               { fmt.Println("dvd player: powering on") }
func (d *DVDPlayer) Off()              { fmt.Println("dvd player: powering off") }
func (d *DVDPlayer) Play(movie string) { fmt.Println("dvd player: playing movie " + movie) }
func (d *DVDPlayer) Stop(movie string) { fmt.Println("dvd player: stopped ") }

type Projector struct{}

func (p *Projector) On()  { fmt.Println("projector: warming uo") }
func (p *Projector) Off() { fmt.Println("projector: cooling down") }
func (p *Projector) WideScreenMode() {
	fmt.Println("Projector: Widescreen mode enabled.")
}

type SmartLights struct{}

func (s *SmartLights) Dim(level int) {
	fmt.Printf("Lights: Dimmed to %d%%.\n", level)
}

func (s *SmartLights) On() {
	fmt.Println("Lights: Full brightness.")
}

type StreamingService struct{}

func (s *StreamingService) Connect() {
	fmt.Println("Streaming: Connected to service.")
}

func (s *StreamingService) Disconnect() {
	fmt.Println("Streaming: Disconnected.")
}

func (s *StreamingService) Stream(movie string) {
	fmt.Println("Streaming: Now streaming '" + movie + "'.")
}

// facade

type HomeTheaterFacade struct {
	amp       *Amplifier
	dvd       *DVDPlayer
	projector *Projector
	lights    *SmartLights
	streaming *StreamingService
}

func NewHomeTheaterFacade(
	amp *Amplifier,
	dvd *DVDPlayer,
	projector *Projector,
	lights *SmartLights,
	streaming *StreamingService,

) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		amp:       amp,
		dvd:       dvd,
		projector: projector,
		lights:    lights,
		streaming: streaming,
	}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("----- preaparing to watch---" + movie + "---")
	h.lights.Dim(10)
	h.projector.On()
	h.projector.WideScreenMode()
	h.amp.On()
	h.amp.SetVolume(20)
	h.streaming.Connect()
	h.streaming.Stream(movie)
	fmt.Println("--- enjoy the movie----")
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("--- shutting down home theaterr ---")
	h.streaming.Disconnect()
	h.amp.Off()
	h.projector.Off()
	h.lights.On()
	fmt.Println("---- home theater off")
}
