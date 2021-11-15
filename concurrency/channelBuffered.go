package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Testing3() {
	var wg sync.WaitGroup

	//we make a buffered channel with only two [][] ways, if we change buffer size, we have more capacity.
	ch := make(chan int, 3)
	for i := 0; i < 10; i++ {
		ch <- i //occupying one way
		wg.Add(1)
		go sendGoroutineBuff(i, &wg, ch)
	}
	defer wg.Wait()
}

func sendGoroutineBuff(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("pid #%d\n", i)
	time.Sleep(3 * time.Second)
	fmt.Printf("pid done #%d\n", i)
	<-c // freeing up occupied space
}
