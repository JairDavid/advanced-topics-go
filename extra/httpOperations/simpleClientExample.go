package httpOperations

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func mymain() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	json, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(json))
}
