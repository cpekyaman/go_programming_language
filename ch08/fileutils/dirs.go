package fileutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// sema is a counting semaphore for limiting concurrency in dirents.
var sema chan struct{}
var cancelled func() bool

func Prepare(done <-chan bool, parallels int) {
	sema = make(chan struct{}, parallels)
	cancelled = func() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	}
}

func WalkDir(dir string, n *sync.WaitGroup, sizes chan<- int64) {
	defer n.Done()

	if cancelled() {
		return
	}

	for _, entry := range dirEntries(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			WalkDir(subdir, n, sizes)
		} else {
			sizes <- entry.Size()
		}
	}
}

func dirEntries(dir string) []os.FileInfo {
	if cancelled() {
		return nil
	}

	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't read dir %s : %v", dir, err)
		return nil
	}
	return files
}
