package main

import (
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
	ShortHelp() string
	Help() string
}

var commands = map[string]command{
	"delete": &deleteCmd{},
	"copy":   &copyCmd{},
	"put":    &putCmd{},
	"insert": &insertCmd{},

	"get":   &getCmd{},
	"range": &rangeCmd{},

	"help": &helpCmd{},
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
		helps = append(helps, help{name: n, text: cmd.ShortHelp()})
	}
	sort.Sort(helps)

	i := 0
	for _, help := range helps {
		fmt.Fprintf(os.Stderr, "  %-*s - %s\n", max, help.name, help.text)
		i++
	}
}
