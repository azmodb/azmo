package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

type rangeCmd struct{}

func (c rangeCmd) Run(ctx context.Context, db *client.DB, args []string) error {
	if len(args) < 2 {
		return errors.New("range: requires at least 2 argument\n")
	}

	var err error
	var rev int64
	if len(args) >= 3 {
		rev, err = strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			return err
		}
	}

	from, to := []byte(args[0]), []byte(args[1])
	pairs, err := db.Range(ctx, client.NewRange(from, to), rev)
	if err != nil {
		return err
	}

	for _, pair := range pairs {
		fmt.Printf("range key:%q revisions:%v revision:%d\n%s\n",
			pair.Key(), pair.Revisions(), pair.Revision(), pair.Value())
	}
	return nil
}

func (c rangeCmd) Name() string { return "range" }
func (c rangeCmd) Args() string { return "from to [rev]" }
func (c rangeCmd) Help() string { return "TODO" }
