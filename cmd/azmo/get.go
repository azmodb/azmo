package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

type getCmd struct{}

func (c getCmd) Run(ctx context.Context, db *client.DB, args []string) (err error) {
	if len(args) < 1 {
		return errors.New("get: requires at least 1 argument")
	}

	var rev int64
	if len(args) >= 2 {
		rev, err = strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return err
		}
	}

	pair, err := db.Get(ctx, []byte(args[0]), rev)
	if err != nil {
		return err
	}

	fmt.Printf("key:%q revisions:%v revision:%d\n%s\n",
		pair.Key(), pair.Revisions(), pair.Revision(), pair.Value())
	return err
}

func (c getCmd) Name() string { return "get" }
func (c getCmd) Help() string { return "TODO" }
