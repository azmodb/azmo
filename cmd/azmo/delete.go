package main

import (
	"errors"
	"fmt"

	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

type deleteCmd struct{}

func (c deleteCmd) Run(ctx context.Context, db *client.DB, args []string) error {
	if len(args) < 1 {
		return errors.New("delete: requires 1 argument")
	}

	key := []byte(args[0])
	res, err := db.Delete(ctx, key)
	if err != nil {
		return err
	}
	fmt.Printf("delete key:%q revisions:%v revision:%d\n", key, res.Revisions(), res.Revision())
	return nil
}

func (c deleteCmd) Name() string      { return "delete" }
func (c deleteCmd) Args() string      { return "key" }
func (c deleteCmd) ShortHelp() string { return "deletes the value for the given key" }
func (c deleteCmd) Help() string      { return "TODO" }
