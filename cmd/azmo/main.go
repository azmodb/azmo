package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

type command interface {
	Run(context.Context, *client.DB, []string) error
	Name() string
	Help() string
}

var commands = map[string]command{
	"delete": &deleteCmd{},
	"copy":   &copyCmd{},
	"put":    &putCmd{},
	"get":    &getCmd{},
}

func main() {
	var addr = flag.String("addr", "localhost:7979", "database service network address")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] command arguments...\n", os.Args[0])
		fmt.Fprint(os.Stderr, usageMsg)
		fmt.Fprintf(os.Stderr, "\nCommands:\n")
		for _, cmd := range commands {
			fmt.Fprintf(os.Stderr, "  %s - %s\n", cmd.Name(), cmd.Help())
		}
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
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

	db, err := client.Dial(*addr, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "database %q: %v\n", addr, err)
		os.Exit(1)
	}
	defer db.Close()

	ctx := context.TODO()
	if err = cmd.Run(ctx, db, args); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s: %v\n", name, err)
		os.Exit(1)
	}
}

const usageMsg = ``
