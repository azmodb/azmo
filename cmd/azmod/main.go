package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/azmodb/azmo/server"
)

const logFlags = log.Ldate | log.Ltime | log.Lmicroseconds

func main() {
	var (
		addr    = flag.String("addr", "localhost:7979", "network listen address")
		network = flag.String("net", "tcp", "stream-oriented network")
		debug   = flag.Bool("debug", false, "write debug messages to stderr")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprint(os.Stderr, usageMsg)
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()

	listener, err := net.Listen(*network, *addr)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer listener.Close()

	donec := make(chan error, 1)
	go func(donec chan<- error) {
		if *debug {
			logger := log.New(os.Stderr, "", logFlags)
			donec <- server.Listen(listener, server.WithLogger(logger))
		} else {
			donec <- server.Listen(listener)
		}
		close(donec)
	}(donec)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	select {
	case err := <-donec:
		if err != nil {
			log.Fatal(err)
		}
	case <-sigc:
		os.Exit(9)
	}
	os.Exit(0)
}

const usageMsg = ``
