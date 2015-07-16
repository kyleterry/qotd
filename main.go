package main

import (
	"log"
	"net"
)

func main() {
	sock, err := net.Listen("tcp", "0.0.0.0:17")
	if err != nil {
		log.Fatal(err)
	}

	defer sock.Close()
	for {
		conn, err := sock.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}

// Goroutine
func handle(conn net.Conn) {
	conn.Write([]byte("some quote"))
	conn.Close()
}
