package server

import (
	"fmt"
	"log"
	"net"
)

func TcpServe(address *string, port *int, handler func(conn net.Conn)) {
	Serve("tcp", address, port, handler)
}

func Serve(protocol string, address *string, port *int, handler func(conn net.Conn)) {
	listener, err := net.Listen(protocol, fmt.Sprintf("%s:%d", *address, *port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handler(conn)
	}
}
