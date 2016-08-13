package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

type command interface {
	Run(context.Context, *client.DB, []string) error
	Name() string
	Args() string
	Help() string
}

var commands = map[string]command{
	"delete": &deleteCmd{},
	"copy":   &copyCmd{},
	"put":    &putCmd{},
	"insert": &insertCmd{},
	"get":    &getCmd{},
	"range":  &rangeCmd{},
}

type help struct {
	name string
	text string
}

type helps []help

func (p helps) Less(i, j int) bool { return p[i].name < p[j].name }
func (p helps) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p helps) Len() int           { return len(p) }

func printDefaults(commands map[string]command) {
	helps := make(helps, 0, len(commands))
	max := 0
	for name, cmd := range commands {
		n := name + " " + cmd.Args()
		if len(n) > max {
			max = len(n)
		}
		helps = append(helps, help{name: n, text: cmd.Help()})
	}
	sort.Sort(helps)

	i := 0
	for _, help := range helps {
		fmt.Fprintf(os.Stderr, "  %-*s - %s\n", max, help.name, help.text)
		i++
	}
}

func main() {
	var addr = flag.String("addr", "localhost:7979", "database service network address")
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
