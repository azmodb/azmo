package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/azmodb/azmo/client"
	"golang.org/x/net/context"
)

var (
	getUsageMsg = `
Get retrieves the value for a key at revision rev. If rev <= 0 get
returns the current value for a key.
`

	getCmd = command{
		Help:      getUsageMsg,
		ShortHelp: "retrieves the value for a key",
		Name:      "get",
		Args:      "[options] key [rev]",
		Run: func(ctx context.Context, db *client.DB, args []string) (err error) {
			set := flag.NewFlagSet("get", flag.ExitOnError)
			set.Usage = func() {
				fmt.Fprintf(stderr, "Usage: %s get [options] key [rev]\n", self)
				fmt.Fprint(stderr, getUsageMsg)
				fmt.Fprintf(stderr, "\nOptions:\n")
				set.PrintDefaults()
				os.Exit(2)
			}

			hist := set.Bool("hist", false, "result includes revision history")
			if err := set.Parse(args); err != nil {
				return err
			}
			args = set.Args()
			if len(args) <= 0 {
				set.Usage()
			}

			var rev int64
			if len(args) >= 2 {
				rev, err = strconv.ParseInt(args[1], 10, 64)
				if err != nil {
					return err
				}
			}

			r := client.NewRange([]byte(args[0]), nil, rev, *hist)
			defer r.Close()

			ev, err := db.Get(ctx, r)
			if err != nil {
				return err
			}
			fmt.Println(ev)
			return nil
		},
	}

	rangeUsageMsg = `
Range retrieves the values in the range [from, to] greater than revision
rev. If rev <= range returns the current value for a key.
`

	rangeCmd = command{
		Help:      rangeUsageMsg,
		ShortHelp: "retrieves the values in the range",
		Name:      "range",
		Args:      "[options] from to [rev]",
		Run: func(ctx context.Context, db *client.DB, args []string) (err error) {
			set := flag.NewFlagSet("get", flag.ExitOnError)
			set.Usage = func() {
				fmt.Fprintf(stderr, "Usage: %s range [options] from to [rev]\n", self)
				fmt.Fprint(stderr, rangeUsageMsg)
				fmt.Fprintf(stderr, "\nOptions:\n")
				set.PrintDefaults()
				os.Exit(2)
			}

			hist := set.Bool("hist", false, "result includes revision history")
			if err := set.Parse(args); err != nil {
				return err
			}
			args = set.Args()
			if len(args) < 2 {
				set.Usage()
			}

			var rev int64
			if len(args) > 2 {
				rev, err = strconv.ParseInt(args[2], 10, 64)
				if err != nil {
					return err
				}
			}

			r := client.NewRange([]byte(args[0]), []byte(args[1]), rev, *hist)
			defer r.Close()

			ch, err := db.Range(ctx, r)
			if err != nil {
				return err
			}
			for ev := range ch {
				if ev.Err() != nil {
					return ev.Err()
				}
				if ev.Event != nil {
					fmt.Println(ev.Event)
				}
			}
			return nil
		},
	}
)
