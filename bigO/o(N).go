package bigO

import "fmt"

//function that recieve an array of integers O(N) where N is the number of elements
func elements(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Print(i) //constant
	}
}
