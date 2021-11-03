package concurrency

import (
	"fmt"
	"net/http"
)

func Available(URL string, chanel chan interface{}) {
	_, err := http.Get(URL)
	if err != nil {
		chanel <- "La URL %s no esta en linea" + URL
	} else {
		chanel <- "La URL %s esta en linea" + URL
	}
}

func SendPeticion() {
	chanel := make(chan interface{})

	urls := []string{
		"https://www.google.com",
		"https://aaaasd",
		"https://www.youtube.com",
		"https://go.dev",
	}
	for _, item := range urls {
		go Available(item, chanel)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-chanel)
	}

}
