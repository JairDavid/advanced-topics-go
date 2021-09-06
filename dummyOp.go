package main

import "fmt"

func nose(array []int, i int) {

	array = append(array[:i], array[i+1:]...)
}

func lop() {
	var array = []int{1, 2, 3, 4}
	fmt.Print(array)
	nose(array, 3)
}
