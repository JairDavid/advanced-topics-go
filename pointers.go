package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Persona struct {
	name     string
	lastname string
	age      uint8
	address  string
	phone    uint16
}

func (p *Persona) showData() {
	fmt.Println("Name: ", p.name)
	fmt.Println("lastname: ", p.lastname)
	fmt.Println("age: ", p.age)
	fmt.Println("address: ", p.address)
	fmt.Println("phone: ", p.phone)
}

func modifyData(p *Persona) {

}

func newRegister(array *[]Persona) {

}

func main() {
	var registers []Persona
	exit := false
	opt := 0
	fmt.Print(registers)
	for !exit {
		fmt.Println("--------- Register people data ---------")
		fmt.Println("1) New register")
		fmt.Println("2) Edit register")
		fmt.Println("3) Delete register")
		fmt.Println("4) Query All")
		fmt.Println("5 exit")
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			newRegister(&registers)
			break
		case 2:
			break
		case 3:
			break
		case 4:
			break
		case 5:
			exit = true
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			break
		}
	}
}
