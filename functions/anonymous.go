package functions

import "fmt"

func StartAnonymous() {
	func() {
		fmt.Println("anonymous fn")
	}()

	fn := func(a string) {
		fmt.Println(a)
	}
	fn("another anonymous fn")
}
