package main

import (
	"fmt"
)

type Person struct {
}

type methods interface {
	sayHello([]byte) (int, error)
}

func (p Person) sayHello(greet []byte) (int, error) {
	value, err := fmt.Println(string(greet))
	return value, err
}

func mainmy() {
	var mtds methods = Person{}

	mtds.sayHello([]byte("hello"))

}
