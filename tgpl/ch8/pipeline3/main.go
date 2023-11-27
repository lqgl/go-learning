package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go counter(naturals)

	// Squares
	go squarer(squares, naturals)

	// Printer
	printer(squares)

}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(out <-chan int) {
	for x := range out {
		fmt.Println(x)
	}
}
