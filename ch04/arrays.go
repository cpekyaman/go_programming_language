package main

import (
	"fmt"
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	var q = [3]int{1, 2, 3}

	for i, v := range q {
		fmt.Printf("%d - %v\n", i, v)
	}

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		TRY
	)

	c := [...]string{USD: "$", EUR: "E"}

	for i, v := range c {
		fmt.Printf("%d - %v\n", i, v)
	}
}
