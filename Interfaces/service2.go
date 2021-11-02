package interfaces

import "fmt"

type Service2 struct {
}

func (s *Service2) Create() {
	fmt.Println("Create from service 2")
}

func (s *Service2) Update() {
	fmt.Println("Update from service 2")
}
