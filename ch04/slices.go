package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	lrotate(s, 2)
	fmt.Println(s)

	s2 := []int{1, 2, 3, 4, 5, 6}
	s2 = remove(s2, 3)
	fmt.Println(s2)
}

func lrotate(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func remove(s []int, n int) []int {
	copy(s[n:], s[n+1:])
	return s[:len(s)-1]
}
