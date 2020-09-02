package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("argument %d = %s\n", i, arg)
	}

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}
