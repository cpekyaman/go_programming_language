package main

import (
	"fmt"

	"github.com/cpekyaman/go_programming_language/ch06/coord"
)

func main() {
	p := coord.Point{X: 1, Y: 2}
	q := coord.Point{X: 4, Y: 6}

	fmt.Println(p.Distance(q))
}
