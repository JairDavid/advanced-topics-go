package concurrency

import "fmt"

//<-chan arrow on left side for only reading
//chan<- arrow on right side for only writting
//close(chan) for close that channel
//https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3

func RandomNumber(rdn chan<- int) {
	for i := 0; i < 10; i++ {
		rdn <- i
	}
	close(rdn)
}

func PowThatNumber(in <-chan int, out chan<- int) {
	for value := range in {
		out <- value * 120
	}
	close(out)
}

func SeeData(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func Testing4() {
	numbers := make(chan int)
	results := make(chan int)

	go RandomNumber(numbers)
	go PowThatNumber(numbers, results)
	SeeData(results)

}
