package concurrencygobook

import (
	"fmt"
)

func RaceCondition() {
	var data int

	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("the value is 0")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
}
