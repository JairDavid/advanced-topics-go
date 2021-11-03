package concurrency

import (
	"fmt"
	"net/http"
)

func Available(URL string) {
	_, err := http.Get(URL)
	if err != nil {
		fmt.Println("La URL %s no esta en linea", URL)
	} else {
		fmt.Println("La URL %s esta en linea", URL)
	}
}

func SendPeticion() {
	urls := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://aaaasd",
		"https://go.dev",
	}
	for _, item := range urls {
		Available(item)
	}
}
