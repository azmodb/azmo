package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/azmodb/azmo"
	pb "github.com/azmodb/azmo/azmopb"
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

func del(ctx context.Context, d *dialer, enc azmo.Encoder, args []string) error {
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

	return c.Delete(ctx, enc, req)
}
