package main

import "fmt"

type Device interface {
	TurnOn()
	TurnOff()
}

type Lamp struct {
	isOn bool
}

func (l *Lamp) TurnOn() {
	if !l.isOn {
		l.isOn = true
		fmt.Println("The lamp is on")
	}
}

func (l *Lamp) TurnOff() {
	if l.isOn {
		l.isOn = false
		fmt.Println("The lamp is off")
	}
}

type Computer struct {
	isOn bool
}

func (c *Computer) TurnOn() {
	if !c.isOn {
		c.isOn = true
		fmt.Println("The computer is on")
	}
}

func (c *Computer) TurnOff() {
	if c.isOn {
		c.isOn = false
		fmt.Println("The computer is off")
	}
}

func TurnOnAll(devices []Device) {
	for _, d := range devices {
		d.TurnOn()
	}
}

func TurnOffAll(devices []Device) {
	for _, d := range devices {
		d.TurnOff()
	}
}

func main() {
	lamp := &Lamp{isOn: false}
	computer := &Computer{isOn: false}

	devices := []Device{lamp, computer}
	TurnOnAll(devices)
	TurnOffAll(devices)
}
