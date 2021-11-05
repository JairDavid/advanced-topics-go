package functions

import "fmt"

func names(listNames ...string) {
	for _, item := range listNames {
		fmt.Println(item)
	}
}

func Start() {
	names("name_1", "name_2", "name_3", "name_4", "name_5")
}
