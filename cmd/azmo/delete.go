package main

import (
	"errors"
	"fmt"

	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

type deleteCmd struct{}

func (c deleteCmd) Run(ctx context.Context, db *client.DB, args []string) (err error) {
	if len(args) < 1 {
		return errors.New("delete: requires 1 argument")
	}

	key := []byte(args[0])
	batch := client.NewTxnBatch()
	batch.Delete(key)

	result, err := db.Apply(ctx, batch)
	if err != nil {
		return err
	}

	for key, num, rev, ok := result.Next(); ok; {
		fmt.Printf("key:%q txnid:%d revision:%d\n", key, num, rev)
		key, num, rev, ok = result.Next()
	}
	return err
}

func (c deleteCmd) Name() string { return "delete" }
func (c deleteCmd) Help() string { return "TODO" }
