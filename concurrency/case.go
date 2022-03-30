package concurrency

import (
	"fmt"
	"time"
)

func Everyms(ch chan<- string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 500)
		ch <- "500ms"
	}
	close(ch)
}

func EverySec(ch chan<- string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 2)
		ch <- "2 seconds"
	}
	close(ch)
}

func Testing6() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go Everyms(ch1)
	go EverySec(ch2)

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
