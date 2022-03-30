package functions

import "fmt"

func names(listNames ...string) {
	for _, item := range listNames {
		fmt.Println(item)
	}
}

func sum(a, b, c, d, e, f int) int {
	return a + b + c + d + e + f
}

func Start() {
	fmt.Print(sum(1, 2, 3, 4, 5, 6))
	names("name_1", "name_2", "name_3", "name_4", "name_5")
}
