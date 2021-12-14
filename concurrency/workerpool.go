package concurrency

import "fmt"

func Worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d: calculating job %d \n", id, job)
		fib := Fibbonacci(job)
		result <- fib
	}
}

func Fibbonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibbonacci(n-1) + Fibbonacci(n-2)
}

func Launch() {
	tasks := []int{25, 125, 342, 443216, 5765, 765, 7978, 788, 569, 310, 1134, 212, 1233, 414, 1235, 136, 173, 1823, 192, 203, 231, 22}
	jobs := make(chan int, len(tasks))
	result := make(chan int, len(tasks))

	for i := 0; i < 3; i++ {
		go Worker(i, jobs, result)
	}

	for task := range tasks {
		jobs <- task
	}

	close(jobs)

	for a := 1; a < len(tasks); a++ {
		<-result
	}
}
