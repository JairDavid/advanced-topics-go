package main

import "fmt"

type Figura interface {
	area() float64
}

type Cuadrado struct {
	lado float64
}

type Rectangulo struct {
	base   float64
	altura float64
}

func (c Cuadrado) area() float64 {
	return c.lado * c.lado
}

func (r Rectangulo) area() float64 {
	return r.base * r.altura
}

func calcularCualquierArea(f Figura) float64 {
	return f.area()
}

func faces() {
	fmt.Println("output")
	f1 := Cuadrado{10}
	f2 := Rectangulo{12, 12}
	figuras := []Figura{f1, f2}

	for _, item := range figuras {
		fmt.Println(calcularCualquierArea(item))
	}
}
