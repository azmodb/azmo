package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/azmodb/azmo/server"
)

func main() {
	var addr = flag.String("addr", "localhost:7979", "network listen address")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprint(os.Stderr, usageMsg)
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()

	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Fatal(server.Listen(listener))
}

const usageMsg = ``
