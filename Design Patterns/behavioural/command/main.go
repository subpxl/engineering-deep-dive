package main

func main() {

	// client code
	tv := &TV{}
	onCommand := &OnCommand{device: tv}
	offCommand := &OffCommand{device: tv}
	onButton := &Button{command: onCommand}
	offButton := &Button{command: offCommand}
	onButton.Press()
	onButton.Undo()
	offButton.Press()
	offButton.Undo()

	light := &Light{}
	onCommandLight := &OnCommand{device: light}
	offCommandLight := &OffCommand{device: light}
	onButtonLight := &Button{command: onCommandLight}
	offButtonLight := &Button{command: offCommandLight}
	onButtonLight.Press()
	offButtonLight.Press()
	onButtonLight.Undo()
}

// invoker
type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}

func (b *Button) Undo() {
	b.command.Undo()
}

// command
type Command interface {
	Execute()
	Undo()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) Execute() {
	c.device.On()
}

func (c *OnCommand) Undo() {
	c.device.Off()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) Execute() {
	c.device.Off()
}

func (c *OffCommand) Undo() {
	c.device.On()
}

// reciever
type Device interface {
	On()
	Off()
}

type TV struct {
	isRunning bool
}

func (t *TV) On() {
	t.isRunning = true
	println("TV is ON")
}

func (t *TV) Off() {
	t.isRunning = false
	println("TV is OFF")
}

//

type Light struct {
	isRunning bool
}

func (t *Light) On() {
	t.isRunning = true
	println("Light is ON")
}

func (t *Light) Off() {
	t.isRunning = false
	println("Light is OFF")
}
