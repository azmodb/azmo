package main

import (
	"flag"
	"log"
	"net"

	"github.com/azmodb/azmo/server"
)

func main() {
	var addr = flag.String("addr", "localhost:7979", "network losten address")
	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Fatal(server.Listen(listener))
}
