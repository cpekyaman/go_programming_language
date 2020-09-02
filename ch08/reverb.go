package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/cpekyaman/go_programming_language/ch08/server"
)

var address = flag.String("h", "localhost", "host")
var port = flag.Int("p", 3000, "port")

func main() {
	flag.Parse()
	server.TcpServe(address, port, handle)
}

func echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
}

func handle(conn net.Conn) {
	defer conn.Close()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), 1*time.Second)
	}
}
