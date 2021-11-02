package main

import "paquetes/interfaces"

func main() {
	var s1 interfaces.Service1
	var s2 interfaces.Service2

	interfaces.Repository.Create(&s1)
	interfaces.Repository.Create(&s2)

}
