package main

import (
	"encoding/json"
	"encoding/xml"
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

type jsonNotifier struct{}

func (n jsonNotifier) Notify(ev *azmo.Event) error {
	return json.NewEncoder(os.Stdout).Encode(ev)
}

type xmlNotifier struct{}

func (n xmlNotifier) Notify(ev *azmo.Event) error {
	return xml.NewEncoder(os.Stdout).Encode(ev)
}

type fmtNotifier struct{}

func (n fmtNotifier) Notify(ev *azmo.Event) error {
	_, err := fmt.Println(ev)
	return err
}

func watch(ctx context.Context, d *dialer, args []string) error {
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

	if *jsonFmt {
		return c.Watch(ctx, req, jsonNotifier{})
	}
	return c.Watch(ctx, req, fmtNotifier{})
}
