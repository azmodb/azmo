package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

var (
	addr    = flag.String("addr", "localhost:7979", "service network address")
	timeout = flag.Duration("timeout", 0, "dialing timeout")
	self    string
	stderr  = os.Stderr
)

func usage() {
	fmt.Fprintf(stderr, "Usage: %s [options] command [options][arguments]...\n",
		self)
	fmt.Fprint(stderr, usageMsg)
	fmt.Fprintf(stderr, "\nOptions:\n\n")
	flag.PrintDefaults()
	fmt.Fprintf(stderr, "\nCommands:\n\n")
	printDefaults()
	fmt.Fprintf(stderr, "%s\n", helpMsg)
	os.Exit(2)
}

func init() { self = os.Args[0] }

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		usage()
	}

	name, args := args[0], args[1:]
	cmd, found := commands[name]
	if !found {
		fmt.Fprintf(stderr, "unknown command %q\n", name)
		usage()
	}

	db, err := client.Dial(*addr, *timeout)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	donec := make(chan error, 1)
	go func(donec chan<- error) {
		donec <- cmd.Run(ctx, db, args)
		close(donec)
	}(donec)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	select {
	case err := <-donec:
		if err != nil {
			log.Fatalf("%v", err)
		}
	case <-sigc:
		cancel()
	}
}

const usageMsg = `
Azmo is a command line client for AzmoDB (https://github.com/azmodb/azmo).
`
