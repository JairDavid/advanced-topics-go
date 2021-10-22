package main

import "fmt"

type testRepo struct {
	text string
}

func (tsr *testRepo) Create(name string) (testRepo, error) {
	fmt.Printf(name)
	return testRepo{}, nil
}

func main() {
	myrepo := new(testRepo)
	typo, err := myrepo.Create("my text")
	if err != nil {
		fmt.Printf("regreso nil")
	}
	fmt.Print(typo)
}
