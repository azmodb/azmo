package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/azmodb/exp/azmo"
	"golang.org/x/net/context"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] command [options][arguments]...\n", self)
	fmt.Fprint(os.Stderr, usageMsg)
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nCommands:\n")
	printDefaults()
	fmt.Fprintf(os.Stderr, "%s\n", helpMsg)
	os.Exit(2)
}

var (
	addr    = flag.String("addr", "localhost:7979", "service network address")
	timeout = flag.Duration("timeout", 0, "dialing timeout")
	jsonFmt = flag.Bool("json", false, "write json encoded output to stdout")
	xmlFmt  = flag.Bool("xml", false, "write xml encoded output to stdout")
	self    string
)

func init() { self = os.Args[0] }

func encode(v interface{}) error {
	if *jsonFmt {
		return json.NewEncoder(os.Stdout).Encode(v)
	}
	if *xmlFmt {
		return xml.NewEncoder(os.Stdout).Encode(v)
	}

	switch t := v.(type) {
	case []*azmo.Event:
		for _, ev := range t {
			fmt.Println(ev)
		}
	case *azmo.Event:
		fmt.Println(t)
	default:
		log.Panicf("cannot encode %T", v)
	}
	return nil
}

func main() {
	flag.Usage = usage
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	args := flag.Args()
	if len(args) <= 0 {
		usage()
	}

	name, args := args[0], args[1:]
	cmd, found := commands[name]
	if !found {
		fmt.Fprintf(os.Stderr, "%s: unknown command %q\n", self, name)
		usage()
	}
	if name == "help" {
		if err := helpCmd.Run(nil, nil, args); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}
	if name == "version" {
		if err := versionCmd.Run(nil, nil, nil); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	ctx, cancel := context.WithCancel(context.Background())
	errc := make(chan error)
	dialer := &dialer{addr: *addr, timeout: *timeout}
	go func() { errc <- cmd.Run(ctx, dialer, args) }()

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	select {
	case err := <-errc:
		if err != nil {
			log.Fatal(err)
		}
	case <-sigc:
		cancel()
	}
}

const usageMsg = `
Azmo is a command line client for AzmoDB (https://github.com/azmodb/azmo).
`
