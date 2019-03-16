package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/azmodb/bricktop/client"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] cmd [option] args...\n", self)
	fmt.Fprint(os.Stderr, usageMsg)
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
	//fmt.Fprintf(os.Stderr, "\nCommands:\n")
	os.Exit(2)
}

var (
	addr    = flag.String("addr", "localhost:5640", "service network address")
	network = flag.String("network", "tcp", "connect on the named network")

	self = string(os.Args[0])
)

const usageMsg = `
9P2000.L client that can access a single file on a 9P2000 server. It can
be useful for manual interaction with a 9P2000.L server or for accessing
simple 9P2000.L services from within scripts.
`

func main() {
	flag.Usage = usage
	flag.Parse()
	//if flag.NArg() < 1 {
	//flag.Usage()
	//}

	ctx, cancel := context.WithCancel(context.Background())
	c, err := client.Dial(ctx, *network, *addr)
	if err != nil {
		log.Fatalf("client connection: %s", err)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)
	select {
	case <-sigc:
		cancel()
	case <-ctx.Done():
		// nothing
	}

	if err = c.Close(); err != nil {
		log.Fatalf("client connection: %s", err)
	}
}
