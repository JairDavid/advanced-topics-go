package pointers

import "fmt"

type Object1 struct {
	name *string
}

func Start() {
	a := &Object1{}

	x := "jair"

	a.name = &x

	fmt.Println(*a.name)
}
