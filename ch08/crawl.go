package main

import (
	"os"

	"github.com/cpekyaman/go_programming_language/ch05/netutils"
)

// solve not-terminating main problem
func main() {
	worklist := make(chan []string)
	linksToProcess := make(chan string)

	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range linksToProcess {
				extracted := netutils.Crawl(link)
				go func() { worklist <- extracted }()
			}
		}()
	}

	processed := make(map[string]bool)
	for list := range worklist {
		list := <-worklist
		for _, link := range list {
			if !processed[link] {
				processed[link] = true
				linksToProcess <- link
			}
		}
	}
}
