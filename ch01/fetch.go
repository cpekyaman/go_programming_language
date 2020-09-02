package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	prefix = "http://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}

		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			fmt.Fprintf(os.Stderr, "fetch Bad Status %d\n", resp.StatusCode)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch Read Body %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}
