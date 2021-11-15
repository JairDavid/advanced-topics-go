package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Testing2() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go sendGoroutine(i, &wg)
	}
	defer wg.Wait()
}

func sendGoroutine(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("pid #%d\n", i)
	time.Sleep(6 * time.Second)
	fmt.Printf("pid done: #%d\n", i)

}
