package main

import "fmt"

type Socket interface {
	turnOn(bool)
	turnOff(bool)
	getStatus() bool
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

type Blender struct {
	Objbase
	rpm int
}

func (o *Objbase) turnOn(socket bool) {
	o.on = socket
}

func (o *Objbase) turnOff(socket bool) {
	o.on = socket
}

func (o *Objbase) getStatus() bool {
	return o.on
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

	fmt.Print(stereo.getStatus())
	fmt.Print(coffeMaker.getStatus())
	fmt.Print(teapot.getStatus())
	fmt.Print(blender.getStatus())

}
