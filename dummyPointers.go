package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type Persona struct {
	name     string
	lastname string
	age      uint8
	address  string
	phone    uint64
}

func (p *Persona) showData() {
	fmt.Println("Name: ", p.name)
	fmt.Println("lastname: ", p.lastname)
	fmt.Println("age: ", p.age)
	fmt.Println("address: ", p.address)
	fmt.Println("phone: ", p.phone)
}

func modifyData(array *[]Persona) {
	var id int
	QueryAll(array)
	fmt.Println("Seleciona un ID.")
	fmt.Scanln(&id)
}

func newRegister(array *[]Persona) {
	var (
		name     string
		lastname string
		age      uint8
		address  string
		phone    uint64
	)
	in := bufio.NewReader(os.Stdin)

	fmt.Println("write your name: ")
	name, _ = in.ReadString('\n')
	fmt.Println("write your lastname: ")
	lastname, _ = in.ReadString('\n')
	fmt.Println("write your age: ")
	fmt.Scanln(&age)
	fmt.Println("write your address: ")
	address, _ = in.ReadString('\n')
	fmt.Println("write your phone: ")
	fmt.Scanln(&phone)
	persona := Persona{name, lastname, age, address, phone}
	*array = append(*array, persona)
}

func QueryAll(array *[]Persona) {
	fmt.Println("----- People List -----")
	if len(*array) == 0 {
		fmt.Println("there's no people registered yet")
	} else {
		for i, item := range *array {
			fmt.Printf("ID: %d,\n data: %v", i, item)
		}
	}
}

func dumm() {
	var registers []Persona
	exit := false
	opt := 0
	fmt.Print(registers)
	for !exit {
		fmt.Println("--------- Register people data ---------")
		fmt.Println("1) New register")
		fmt.Println("2) Edit register")
		fmt.Println("4) Query All")
		fmt.Println("5) exit")
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			newRegister(&registers)
			break
		case 2:
			modifyData(&registers)
			break
		case 3:
			break
		case 4:
			QueryAll(&registers)
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
