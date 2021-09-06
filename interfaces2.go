package main

import "fmt"

type Socket interface {
	turnOn(socket bool)
	turnOff(socket bool)
}

type Objbase struct {
	on bool
}

type Stereo struct {
	Objbase
	vol int
}

type CoffeMaker struct {
	Objbase
	capacity int
}

type Teapot struct {
	Objbase
	maxTemp int
}

func (o *Objbase) turnOn(socket bool) {
	o.on = socket
}

func (o *Objbase) turnOff(socket bool) {
	o.on = socket
}

type Blender struct {
	Objbase
	rpm int
}

func main() {

	stereo := Stereo{Objbase{false}, 10}
	coffeMaker := CoffeMaker{Objbase{false}, 10}
	teapot := Teapot{Objbase{false}, 120}
	blender := Blender{Objbase{false}, 20}

	fmt.Println(stereo, coffeMaker, teapot, blender)

	stereo.turnOn(true)
	coffeMaker.turnOn(true)
	teapot.turnOn(true)
	blender.turnOn(true)

	fmt.Println(stereo, coffeMaker, teapot, blender)
}
