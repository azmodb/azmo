package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	pb "github.com/azmodb/exp/azmo/azmopb"
	"golang.org/x/net/context"
)

var getCmd = command{
	Help: `
Get retrieves the value for a key at revision. If revision <= 0 it
returns the current value for a key. If equal is true the value
revision must match the supplied rev.
`,
	Short: "retrieves the value for a key",
	Args:  "[options] key [revision]",
	Run:   get,
}

func get(ctx context.Context, d *dialer, args []string) (err error) {
	flags := flag.FlagSet{}
	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s: get [options] key [revision]\n", self)
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flags.PrintDefaults()
		os.Exit(2)
	}
	equal := flags.Bool("equal", false, "revision must equal value revision")
	flags.Parse(args)
	if flags.NArg() < 1 {
		flags.Usage()
	}
	args = flags.Args()
	var rev int64
	if len(args) >= 2 {
		rev, err = strconv.ParseInt(args[1], 10, 0)
		if err != nil {
			return err
		}
	}

	req := &pb.GetRequest{
		Key:       []byte(args[0]),
		Revision:  rev,
		MustEqual: *equal,
	}

	c := d.dial()
	defer c.Close()

	ev, err := c.Get(ctx, req)
	if err != nil {
		return err
	}

	return encode(ev)
}
