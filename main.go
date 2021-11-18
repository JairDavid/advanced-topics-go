package main

import (
	"fmt"
	interfaces "paquetes/Interfaces"
)

func main() {
	var s1 interfaces.Service1
	var s2 interfaces.Service2

	fmt.Println(interfaces.Repository.Create(&s1, "jair 1"))
	fmt.Println(interfaces.Repository.Create(&s2, "jair 2"))
	//interfaces.TestingImplementation()
	//functions.Start()
	// functions.StartAnonymous()

	// errmanagment.OpenFile()

	// fn := functions.Repeat(5)
	// fmt.Println(fn("testing"))
	// count := functions.Counter()
	// fmt.Println(count())
	// fmt.Println(count())

	// concurrency.SendPeticion()

	//concurrency.Testing6()
	//concurrency.Testing5()
	// concurrency.Testing4()
	// concurrency.Testing3()
	// concurrency.Testing2()
	// concurrency.Testing()
}
