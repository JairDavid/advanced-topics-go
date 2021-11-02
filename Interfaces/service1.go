package interfaces

import "fmt"

type Service1 struct {
}

func (s *Service1) Create() {
	fmt.Println("Create from service 1")
}

func (s *Service1) Update() {
	fmt.Println("Update from service 1")
}
