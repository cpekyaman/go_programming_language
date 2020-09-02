package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fn := fibn(n)
	fmt.Printf("\rfibonacci of %d is %d\n", n, fn)

}

func spinner(duration time.Duration) {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(duration)
		}
	}
}

func fibn(x int) int {
	if x < 2 {
		return x
	}
	return fibn(x-1) + fibn(x-2)
}
