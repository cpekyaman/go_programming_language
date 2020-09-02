package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"bob":  25,
		"alex": 30,
		"john": 32,
		"mary": 38,
	}

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	fmt.Println()

	ages["bob"]++

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	fmt.Println()
}
