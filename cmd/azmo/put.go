package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

type putCmd struct{}

func (c putCmd) Run(ctx context.Context, db *client.DB, args []string) (err error) {
	if len(args) < 2 {
		return errors.New("put: requires at least 2 arguments")
	}

	key, value := []byte(args[0]), []byte(args[1])

	batch := client.NewTxnBatch()
	batch.Put(key, value, false)

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

func (c putCmd) Name() string { return "put" }
func (c putCmd) Help() string { return "TODO" }

type copyCmd struct{}

func (c copyCmd) Run(ctx context.Context, db *client.DB, args []string) error {
	if len(args) < 2 {
		return errors.New("copy: requires at least 2 arguments")
	}

	key, path := []byte(args[0]), args[1]
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if fi.IsDir() {
		panic("TODO: handle directories")
	}

	value, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	batch := client.NewTxnBatch()
	batch.Put(key, value, false)

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

func (c copyCmd) Name() string { return "copy" }
func (c copyCmd) Help() string { return "TODO" }
