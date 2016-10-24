package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] command [options][arguments]...\n", self)
	fmt.Fprint(os.Stderr, usageMsg)
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nCommands:\n")
	printDefaults()
	fmt.Fprintf(os.Stderr, "%s\n", helpMsg)
	os.Exit(2)
}

var (
	addr    = flag.String("addr", "localhost:7979", "service network address")
	timeout = flag.Duration("timeout", 0, "dialing timeout")
	self    string
)

func init() { self = os.Args[0] }

func main() {
	flag.Usage = usage
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	args := flag.Args()
	if len(args) <= 0 {
		usage()
	}

	name, args := args[0], args[1:]
	cmd, found := commands[name]
	if !found {
		fmt.Fprintf(os.Stderr, "%s: unknown command %q\n", self, name)
		usage()
	}
	if name == "help" {
		if err := helpCmd.Run(nil, nil, nil, args); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	/*
		c, err := azmo.Dial(*addr, *timeout)
		if err != nil {
			log.Fatalf("dial azmo server: %v", err)
		}
		defer c.Close()
	*/

	ctx, cancel := context.WithCancel(context.Background())
	errc := make(chan error)
	dialer := &dialer{addr: *addr, timeout: *timeout}
	go func() { errc <- cmd.Run(ctx, dialer, fmtEncoder{}, args) }()

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	select {
	case err := <-errc:
		if err != nil {
			log.Fatal(err)
		}
	case <-sigc:
		cancel()
	}
}

const usageMsg = `
Azmo is a command line client for AzmoDB (https://github.com/azmodb/azmo).
`
