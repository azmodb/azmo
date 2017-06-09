package main

import (
	"flag"
	"fmt"
	"os"

	pb "github.com/azmodb/exp/azmo/azmopb"
	"golang.org/x/net/context"
)

var delCmd = command{
	Help: `
Delete removes a key/value pair.
`,
	Short: "removes a key/value pair",
	Args:  "key",
	Run:   put,
}

func del(ctx context.Context, d *dialer, args []string) error {
	flags := flag.FlagSet{}
	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s: delete key\n", self)
		os.Exit(2)
	}
	flags.Parse(args)
	if flags.NArg() != 1 {
		flags.Usage()
	}
	args = flags.Args()

	req := &pb.DeleteRequest{
		Key: []byte(args[0]),
	}

	c := d.dial()
	defer c.Close()

	ev, err := c.Delete(ctx, req)
	if err != nil {
		return err
	}

	return encode(ev)
}
