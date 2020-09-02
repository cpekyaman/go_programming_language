package main

import (
	"fmt"

	"github.com/cpekyaman/go_programming_language/ch02/tempconv"
)

func main() {
	var c tempconv.Celsius = 20
	fmt.Println(tempconv.CToF(c))
}
