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
	putUsageMsg = `
Put sets the value for a key. If the key exists and tombstone is supplied
then its previous versions will be overwritten. If tombstone is false a
new version will be created.
`

	putCmd = command{
		Help:      putUsageMsg,
		ShortHelp: "put sets the value for a key",
		Name:      "put",
		Args:      "[options] key value",
		Run: func(ctx context.Context, db *client.DB, args []string) (err error) {
			set := flag.NewFlagSet("put", flag.ExitOnError)
			set.Usage = func() {
				fmt.Fprintf(stderr, "Usage: %s put [options] key value\n", self)
				fmt.Fprint(stderr, putUsageMsg)
				fmt.Fprintf(stderr, "\nOptions:\n")
				set.PrintDefaults()
				os.Exit(2)
			}

			tombstone := set.Bool("tombstone", false, "delete previous revisions")
			prev := set.Bool("prev", false, "return previous value, if any")

			if err := set.Parse(args); err != nil {
				return err
			}
			args = set.Args()
			if len(args) < 2 {
				set.Usage()
			}

			b := client.NewBatch()
			if !*tombstone {
				b.Insert([]byte(args[0]), []byte(args[1]), *prev)
			} else {
				b.Put([]byte(args[0]), []byte(args[1]), *prev)
			}
			defer b.Close()

			ch, err := db.Apply(ctx, b)
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

	incrementUsageMsg = `
Inc increments the value for a key.
`

	incrementCmd = command{
		Help:      incrementUsageMsg,
		ShortHelp: "inc increments the value for a key",
		Name:      "inc",
		Args:      "[options] key value",
		Run: func(ctx context.Context, db *client.DB, args []string) (err error) {
			set := flag.NewFlagSet("inc", flag.ExitOnError)
			set.Usage = func() {
				fmt.Fprintf(stderr, "Usage: %s inc [options] key value\n", self)
				fmt.Fprint(stderr, incrementUsageMsg)
				fmt.Fprintf(stderr, "\nOptions:\n")
				set.PrintDefaults()
				os.Exit(2)
			}

			prev := set.Bool("prev", false, "return previous value, if any")
			if err := set.Parse(args); err != nil {
				return err
			}
			args = set.Args()
			if len(args) < 2 {
				set.Usage()
			}

			val, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			b := client.NewBatch()
			b.Increment([]byte(args[0]), int64(val), *prev)
			defer b.Close()

			ch, err := db.Apply(ctx, b)
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

	decrementUsageMsg = `
Dec decrements the value for a key.
`

	decrementCmd = command{
		Help:      incrementUsageMsg,
		ShortHelp: "dec decrements the value for a key",
		Name:      "dec",
		Args:      "[options] key value",
		Run: func(ctx context.Context, db *client.DB, args []string) (err error) {
			set := flag.NewFlagSet("dec", flag.ExitOnError)
			set.Usage = func() {
				fmt.Fprintf(stderr, "Usage: %s dec [options] key value\n", self)
				fmt.Fprint(stderr, decrementUsageMsg)
				fmt.Fprintf(stderr, "\nOptions:\n")
				set.PrintDefaults()
				os.Exit(2)
			}

			prev := set.Bool("prev", false, "return previous value, if any")
			if err := set.Parse(args); err != nil {
				return err
			}
			args = set.Args()
			if len(args) < 2 {
				set.Usage()
			}

			val, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			b := client.NewBatch()
			b.Decrement([]byte(args[0]), int64(val), *prev)
			defer b.Close()

			ch, err := db.Apply(ctx, b)
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
