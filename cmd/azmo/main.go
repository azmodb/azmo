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

func main() {
	var (
		addr    = flag.String("addr", "localhost:7979", "database service network address")
		timeout = flag.Duration("timeout", 0, "database dialing timeout")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] command arguments...\n", os.Args[0])
		fmt.Fprint(os.Stderr, usageMsg)
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nCommands:\n")
		printDefaults(commands)
		os.Exit(2)
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
	}

	name, args := args[0], args[1:]
	cmd, found := commands[name]
	if !found {
		fmt.Fprintf(os.Stderr, "unknown command %s\n", name)
		flag.Usage()
	}
	if name == "help" {
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "help: requires 1 argument\n")
			os.Exit(2)
		}

		c, found := commands[args[0]]
		if !found {
			fmt.Fprintf(os.Stderr, "help: command not found\n")
			os.Exit(1)
		}
		fmt.Printf("%s\n", c.Help())
		return
	}

	db, err := client.Dial(*addr, *timeout)
	if err != nil {
		log.Fatal("database %q: %v\n", addr, err)
	}
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	donec := make(chan error, 1)
	go func(ctx context.Context, donec chan<- error) {
		donec <- cmd.Run(ctx, db, args)
		close(donec)
	}(ctx, donec)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	select {
	case err := <-donec:
		if err != nil {
			log.Fatal(err)
		}
	case <-sigc:
		cancel()
	}
}

const usageMsg = ``
