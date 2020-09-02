package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		processFiles(files, counts)
	}

	printLines(counts)
}

func processFiles(files []string, counts map[string]int) {
	for _, fname := range files {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Println(err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}
}

func printLines(counts map[string]int) {
	for line, count := range counts {
		fmt.Printf("%s\t%d\n", line, count)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}
