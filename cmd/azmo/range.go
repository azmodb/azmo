package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/azmodb/azmo"
	pb "github.com/azmodb/azmo/azmopb"
	"golang.org/x/net/context"
)

var (
	rangeCmd = command{
		Help: `
Range iterates over values stored in the database in the range at rev
over the interval [from, to] from left to right. Limit limits the
number of keys returned. If revision <= 0 range gets the keys at the
current revision of the database.
`,
		Short: "range over stored values",
		Args:  "[options] from to [revision]",
		Run:   scan,
	}
	forEachCmd = command{
		Help: `
Range iterates over a;; values stored in the database in the range at
revision from left to right. Limit limits the number of keys returned.
If revision <= 0 range gets the keys at the current revision of the
database.
`,
		Short: "range over all stored values",
		Args:  "[options] [revision]",
		Run:   forEach,
	}
)

func scan(ctx context.Context, d *dialer, enc azmo.Encoder, args []string) (err error) {
	flags := flag.FlagSet{}
	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s: range [options] from to [revision]\n", self)
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flags.PrintDefaults()
		os.Exit(2)
	}
	limit := flags.Int("limit", 0, "maximum range query results")
	flags.Parse(args)
	if flags.NArg() < 2 {
		flags.Usage()
	}
	args = flags.Args()
	var rev int64
	if len(args) >= 3 {
		rev, err = strconv.ParseInt(args[2], 10, 0)
		if err != nil {
			return err
		}
	}

	req := &pb.RangeRequest{
		From:     []byte(args[0]),
		To:       []byte(args[1]),
		Revision: rev,
		Limit:    int32(*limit),
	}

	c := d.dial()
	defer c.Close()

	return c.Range(ctx, enc, req)
}

func forEach(ctx context.Context, d *dialer, enc azmo.Encoder, args []string) (err error) {
	flags := flag.FlagSet{}
	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s: foreach [options] [revision]\n", self)
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flags.PrintDefaults()
		os.Exit(2)
	}
	limit := flags.Int("limit", 0, "maximum range query results")
	flags.Parse(args)
	args = flags.Args()
	var rev int64
	if len(args) >= 1 {
		rev, err = strconv.ParseInt(args[0], 10, 0)
		if err != nil {
			return err
		}
	}

	req := &pb.RangeRequest{
		From:     nil,
		To:       nil,
		Revision: rev,
		Limit:    int32(*limit),
	}

	c := d.dial()
	defer c.Close()

	return c.Range(ctx, enc, req)
}