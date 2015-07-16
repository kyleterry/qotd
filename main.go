package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
)

var quotes []string

func main() {
	//make cache of quotes
	input, err := ioutil.ReadFile("/etc/qotd/quotes.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(input, &quotes); err != nil {
		log.Fatal(err)
	}

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
	l := rand.Perm(len(quotes) - 1)
	conn.Write([]byte(quotes[l[0]]))
	conn.Close()
}
