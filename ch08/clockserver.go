package main

import (
	"flag"
	"io"
	"net"
	"time"

	"github.com/cpekyaman/go_programming_language/ch08/server"
)

var address = flag.String("h", "localhost", "host")
var port = flag.Int("p", 3000, "port")

func main() {
	flag.Parse()

	server.TcpServe(address, port, handle)
}

func handle(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("15:05:04\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
