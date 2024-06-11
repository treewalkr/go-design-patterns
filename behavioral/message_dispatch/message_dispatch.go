package messagedispatch

import (
	"fmt"
)

// Message defines a generic command for smart home devices
type Message struct {
	Action string
	Args   []interface{}
}

// Device is the interface representing any smart home device
type Device interface {
	HandleMessage(msg Message) interface{}
}

// Light is a smart device that can turn on, off, and dim
type Light struct {
	brightness int
}

// HandleMessage processes commands sent to the Light
func (l *Light) HandleMessage(msg Message) interface{} {
	switch msg.Action {
	case "turn_on":
		l.brightness = 100
		return "Light turned on"
	case "turn_off":
		l.brightness = 0
		return "Light turned off"
	case "dim":
		level := msg.Args[0].(int)
		l.brightness = level
		return fmt.Sprintf("Light dimmed to %d%% brightness", level)
	default:
		return "Unknown command for Light"
	}
}

// Fan is a smart device that can turn on, off, and adjust speed
type Fan struct {
	speed int
}

// HandleMessage processes commands sent to the Fan
func (f *Fan) HandleMessage(msg Message) interface{} {
	switch msg.Action {
	case "turn_on":
		f.speed = 1
		return "Fan turned on at speed 1"
	case "turn_off":
		f.speed = 0
		return "Fan turned off"
	case "set_speed":
		speed := msg.Args[0].(int)
		f.speed = speed
		return fmt.Sprintf("Fan speed set to %d", speed)
	default:
		return "Unknown command for Fan"
	}
}

// SendMessage is a helper function to send a message to any device
func SendMessage(device Device, msg Message) interface{} {
	return device.HandleMessage(msg)
}

func Example() {
	// Create smart home devices
	light := &Light{}
	fan := &Fan{}

	// Interact with Light
	fmt.Println(SendMessage(light, Message{Action: "turn_on"}))
	fmt.Println(SendMessage(light, Message{Action: "dim", Args: []interface{}{50}}))
	fmt.Println(SendMessage(light, Message{Action: "turn_off"}))

	// Interact with Fan
	fmt.Println(SendMessage(fan, Message{Action: "turn_on"}))
	fmt.Println(SendMessage(fan, Message{Action: "set_speed", Args: []interface{}{3}}))
	fmt.Println(SendMessage(fan, Message{Action: "turn_off"}))
}
