package main

import (
	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

type helpCmd struct{}

func (c helpCmd) Run(_ context.Context, _ *client.DB, args []string) error {
	return nil
}

func (c helpCmd) Name() string      { return "help" }
func (c helpCmd) Args() string      { return "command" }
func (c helpCmd) ShortHelp() string { return "TOOD" }
func (c helpCmd) Help() string      { return "TODO" }
