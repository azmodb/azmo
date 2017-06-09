package main

import (
	"flag"
	"fmt"
	"os"

	pb "github.com/azmodb/exp/azmo/azmopb"
	"golang.org/x/net/context"
)

var putCmd = command{
	Help: `
Put sets the value for a key. If the key exists and tombstone is true
then its previous versions will be overwritten. Supplied key and
value must remain valid for the life of the database.

It the key exists and the value data type differ, it returns an error.
`,
	Short: "sets the value for a key",
	Args:  "[options] key value",
	Run:   put,
}

func put(ctx context.Context, d *dialer, args []string) error {
	flags := flag.FlagSet{}
	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s: put [options] key value\n", self)
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flags.PrintDefaults()
		os.Exit(2)
	}
	tombstone := flags.Bool("tombstone", false, "overwrite previous values")
	flags.Parse(args)
	if flags.NArg() != 2 {
		flags.Usage()
	}
	args = flags.Args()

	req := &pb.PutRequest{
		Key:       []byte(args[0]),
		Value:     []byte(args[1]),
		Tombstone: *tombstone,
	}

	c := d.dial()
	defer c.Close()

	ev, err := c.Put(ctx, req)
	if err != nil {
		return err
	}

	return encode(ev)
}
