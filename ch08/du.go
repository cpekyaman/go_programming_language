package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/cpekyaman/go_programming_language/ch08/fileutils"
)

var verbose = flag.Bool("v", false, "verbose progress logging")
var parallels = flag.Int("p", 20, "how many parallel dir walks")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var tick chan<- time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// Cancel traversal when input is detected.
	var done = make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// walk directories
	var n sync.WaitGroup
	var nfiles, nbytes int64
	sizes := make(chan int64)

	fileutils.Prepare(done, &parallels)

	for _, root := range roots {
		n.Add(1)
		go fileutils.WalkDir(root, &n, sizes)
	}

	// wait and close sizes channel
	go func() {
		n.Wait()
		close(sizes)
	}()

loop:
	for {
		select {
		case <-done:
			for range sizes {
				// drain channel
			}
			return
		case size, ok := <-sizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printUsage(nfiles, nbytes)
		}
	}

	printUsage(nfiles, nbytes)
}

func printUsage(nfiles, nbytes int64) {
	fmt.Fprintf(os.Stdout, "%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
