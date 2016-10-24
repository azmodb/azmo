package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"text/tabwriter"
	"time"

	"github.com/azmodb/azmo"
	"github.com/azmodb/azmo/build"
	"golang.org/x/net/context"
)

type command struct {
	Run   func(ctx context.Context, d *dialer, args []string) error
	Args  string
	Help  string
	Short string
}

type dialer struct {
	addr    string
	timeout time.Duration
}

func (d *dialer) dial() *azmo.Client {
	c, err := azmo.Dial(d.addr, d.timeout)
	if err != nil {
		log.Fatalf("dialing azmo database server: %v")
	}
	return c
}

var commands = map[string]command{}

func init() {
	commands["foreach"] = forEachCmd
	commands["range"] = rangeCmd
	commands["get"] = getCmd
	commands["watch"] = watchCmd

	commands["put"] = putCmd
	commands["delete"] = delCmd

	commands["version"] = versionCmd
	commands["help"] = helpCmd
}

const helpMsg = `
Use "azmo help [command]" for more information about a command.
`

var (
	helpCmd = command{
		Help:  helpMsg,
		Short: "information about a command",
		Args:  "command",
		Run: func(_ context.Context, d *dialer, args []string) error {
			if len(args) <= 0 {
				fmt.Fprintln(os.Stderr, helpMsg)
				os.Exit(2)
			}

			cmd, found := commands[args[0]]
			if !found {
				return fmt.Errorf("%s: unknown command %q", self, args[0])
			}
			fmt.Println(cmd.Help)
			return nil
		},
	}
	versionCmd = command{
		Help:  "Information about AzmoDB build version",
		Short: "information about version",
		Args:  "",
		Run: func(_ context.Context, d *dialer, args []string) error {
			version()
			return nil
		},
	}
)

func version() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintf(tw, "AzmoDB Version:\t%s\n", build.Version())
	fmt.Fprintf(tw, "ARCH:\t%s\n", runtime.GOARCH)
	fmt.Fprintf(tw, "OS:\t%s\n", runtime.GOOS)
	tw.Flush()
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
		helps = append(helps, help{name: n, text: cmd.Short})
	}
	sort.Sort(helps)

	i := 0
	for _, help := range helps {
		fmt.Fprintf(os.Stderr, "  %-*s - %s\n",
			max, help.name, help.text)
		i++
	}
}
