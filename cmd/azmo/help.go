package main

import (
	"fmt"
	"os"

	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

const helpMsg = `
Use "azmo help [command]" for more information about a command.
`

var helpCmd = command{
	Help:      helpMsg,
	ShortHelp: "information about a command",
	Name:      "help",
	Args:      "[command]",
	Run: func(_ context.Context, _ *client.DB, args []string) error {
		if len(args) <= 0 {
			fmt.Fprintln(stderr, helpMsg)
			os.Exit(2)
		}

		cmd, found := commands[args[0]]
		if !found {
			return fmt.Errorf("command %q not found", args[0])
		}
		fmt.Println(cmd.Help)
		return nil
	},
}
