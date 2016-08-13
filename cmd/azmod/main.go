package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"

	"github.com/azmodb/azmo/server"
)

func main() {
	var (
		debug = flag.Bool("debug", false, "print debug messages to standard output")
		name  = flag.String("name", "", "debug messages prefix")

		addr    = flag.String("addr", "localhost:7979", "network listen address")
		network = flag.String("net", "tcp", "stream-oriented network")
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
			if len(*name) > 0 {
				*name = strings.ToUpper(*name) + " "
			}
			logger := log.New(os.Stdout, *name, log.Ldate|log.Lmicroseconds)
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
	}
}

const usageMsg = ``
