package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

type command struct {
	Run       func(ctx context.Context, c *client.DB, args []string) error
	Name      string
	Args      string
	Help      string
	ShortHelp string
}

var commands = map[string]command{}

func register(cmd command) {
	commands[cmd.Name] = cmd
}

func init() {
	register(incrementCmd)
	register(decrementCmd)
	register(putCmd)
	register(getCmd)
	register(helpCmd)
}

type help struct {
	name string
	text string
}

type helps []help

func (p helps) Less(i, j int) bool { return p[i].name < p[j].name }
func (p helps) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p helps) Len() int           { return len(p) }

func printDefaults() {
	helps := make(helps, 0, len(commands))
	max := 0
	for name, cmd := range commands {
		n := name + " " + cmd.Args
		if len(n) > max {
			max = len(n)
		}
		helps = append(helps, help{name: n, text: cmd.ShortHelp})
	}
	sort.Sort(helps)

	i := 0
	for _, help := range helps {
		fmt.Fprintf(os.Stderr, "  %-*s - %s\n",
			max, help.name, help.text)
		i++
	}
}
