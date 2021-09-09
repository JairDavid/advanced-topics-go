package main

import "fmt"

type Methods interface {
	connect(bool)
	status()
}

type Deltails struct {
	name   string
	model  string
	madeIn string
}

type Television struct {
	Deltails
	on      bool
	channel int
	volume  int
}

type Radio struct {
	Deltails
	on        bool
	frecuency float64
	volume    int
}

type Laptop struct {
	Deltails
	on       bool
	internet bool
}

func (d *Deltails) showDeltails() {
	fmt.Println("Name: ", d.name)
	fmt.Println("Model: ", d.model)
	fmt.Println("Maded in: ", d.madeIn)
}

func (t *Television) connect(ctd bool) {

}
func (t *Radio) connect(ctd bool) {

}
func (t *Laptop) connect(ctd bool) {

}

func maina() {
	fmt.Println("xd")
}
