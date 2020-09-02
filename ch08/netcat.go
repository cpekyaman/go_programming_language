package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/cpekyaman/go_programming_language/ch08/client"
)

var address = flag.String("h", "localhost", "host")
var port = flag.Int("p", 3000, "port")
var interactive = flag.Bool("i", true, "interactive")

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *address, *port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if *interactive {
		go client.MustCopy(os.Stdout, conn)
		client.MustCopy(conn, os.Stdin)
	} else {
		client.MustCopy(os.Stdout, conn)
	}
}
