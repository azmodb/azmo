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

			versions := set.Bool("versions", false, "result includes all revisions")
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

			r := client.NewRange([]byte(args[0]), nil, rev, *versions)
			defer r.Close()

			ev, err := db.Get(ctx, r)
			if err != nil {
				return err
			}
			fmt.Println(ev)
			return nil
		},
	}
)
