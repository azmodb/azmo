package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

type copyCmd struct{}

func copyDir(ctx context.Context, db *client.DB, key []byte, root string) error {
	batch := client.NewTxnBatch()

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)

		value, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		batch.Put([]byte(path), value)
		return nil
	})
	if err != nil {
		return err
	}

	result, err := db.Txn(ctx, batch)
	if err != nil {
		return err
	}

	for _, res := range result {
		fmt.Printf("copy dir: key:%q txnid:%d revision:%d\n",
			res.Key(), res.Num(), res.Revision())
	}
	return err
}

func (c copyCmd) Run(ctx context.Context, db *client.DB, args []string) error {
	if len(args) < 2 {
		return errors.New("copy: requires 2 arguments")
	}

	key, path := []byte(filepath.Clean(args[0])), filepath.Clean(args[1])

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
		return copyDir(ctx, db, key, path)
	}

	value, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	rev, err := db.Put(ctx, key, value)
	if err != nil {
		return err
	}
	fmt.Printf("copy key:%q revision:%d\n", key, rev)
	return nil
}

func (c copyCmd) Name() string { return "copy" }
func (c copyCmd) Args() string { return "key path" }
func (c copyCmd) Help() string { return "TODO" }

type putCmd struct{}

func (c putCmd) Run(ctx context.Context, db *client.DB, args []string) error {
	if len(args) < 2 {
		return errors.New("put: requires 2 arguments")
	}

	key, value := []byte(args[0]), []byte(args[1])
	rev, err := db.Put(ctx, key, value)
	if err != nil {
		return err
	}
	fmt.Printf("put key:%q revision:%d\n", key, rev)
	return nil
}

func (c putCmd) Name() string { return "put" }
func (c putCmd) Args() string { return "key value" }
func (c putCmd) Help() string { return "TODO" }

type insertCmd struct{}

func (c insertCmd) Run(ctx context.Context, db *client.DB, args []string) error {
	if len(args) < 2 {
		return errors.New("insert: requires 2 arguments")
	}

	key, value := []byte(args[0]), []byte(args[1])
	rev, err := db.Insert(ctx, key, value)
	if err != nil {
		return err
	}
	fmt.Printf("insert key:%q revision:%d\n", key, rev)
	return nil
}

func (c insertCmd) Name() string { return "insert" }
func (c insertCmd) Args() string { return "key value" }
func (c insertCmd) Help() string { return "TODO" }
