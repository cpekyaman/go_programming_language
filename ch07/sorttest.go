package main

import (
	"fmt"
	"sort"

	"github.com/cpekyaman/go_programming_language/ch07/sorting"
)

func main() {
	names := []string{"john", "cenk", "alex", "mary", "mike", "ellen"}
	ss := sorting.StringSlice(names)
	sort.Sort(ss)
	fmt.Println(ss)
}
