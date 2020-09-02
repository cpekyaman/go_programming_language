package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	go generator(30, numbers)
	go mapper(numbers, squares)
	reducer(squares)
}

func generator(limit int, out chan<- int) {
	for x := 0; x < limit; x++ {
		out <- x
		time.Sleep(100 * time.Millisecond)
	}
	close(out)
}

func mapper(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
		time.Sleep(100 * time.Millisecond)
	}
	close(out)
}

func reducer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
