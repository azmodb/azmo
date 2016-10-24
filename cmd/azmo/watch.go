package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/azmodb/azmo"
	pb "github.com/azmodb/azmo/azmopb"
	"golang.org/x/net/context"
)

var watchCmd = command{
	Help: `
Watch returns a notifier for a key. If the key does not exist it
returns an error.
`,
	Short: "notifier for a key",
	Args:  "key",
	Run:   watch,
}

func watch(ctx context.Context, d *dialer, enc azmo.Encoder, args []string) error {
	flags := flag.FlagSet{}
	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s: watch key", self)
		os.Exit(2)
	}
	flags.Parse(args)
	if flags.NArg() != 1 {
		flags.Usage()
	}
	args = flags.Args()

	req := &pb.WatchRequest{
		Key: []byte(args[0]),
	}

	c := d.dial()
	defer c.Close()

	return c.Watch(ctx, enc, req)
}
