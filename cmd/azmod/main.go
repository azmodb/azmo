package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"text/tabwriter"

	"github.com/azmodb/azmo"
	"github.com/azmodb/azmo/build"
	"github.com/azmodb/db"
)

//var azmoVersion = fmt.Sprintf("alpha %s/%s %s", runtime.GOOS, runtime.GOARCH, time.Now())

func main() {
	var (
		addr     = flag.String("addr", "localhost:7979", "network listen address")
		network  = flag.String("net", "tcp", "stream-oriented network")
		certFile = flag.String("cert", "", "TLS cert file")
		keyFile  = flag.String("key", "", "TLS key file")
		debug    = flag.Bool("debug", false, "write debug messages to stderr")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [version]\n", os.Args[0])
		fmt.Fprint(os.Stderr, usageMsg)
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
		os.Exit(2)
	}
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	if flag.NArg() == 1 && flag.Arg(0) == "version" {
		version()
		os.Exit(0)
	}

	version()

	listener, err := net.Listen(*network, *addr)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer listener.Close()

	donec := make(chan error, 1)
	go func(donec chan<- error) {
		db := db.New()

		if *debug {
			//logger := log.New(os.Stderr, "", logFlags)
			//donec <- server.Listen(db, listener, server.WithLogger(logger))
		} else {
			donec <- azmo.Listen(db, listener, *certFile, *keyFile)
		}
		close(donec)
	}(donec)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	select {
	case err := <-donec:
		if err != nil {
			log.Fatal(err)
		}
	case <-sigc:
		os.Exit(9)
	}
	os.Exit(0)
}

func version() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintf(tw, "AzmoDB Version:\t%s\n", build.Version())
	fmt.Fprintf(tw, "ARCH:\t%s\n", runtime.GOARCH)
	fmt.Fprintf(tw, "OS:\t%s\n", runtime.GOOS)
	tw.Flush()
}

const usageMsg = `
AzmoDB is an immutable, consistent, in-memory key/value store. AzmoDB
uses an immutable Left-Leaning Red-Black tree (LLRB) internally and
supports snapshotting.

The database provides Atomicity, Consistency and Isolation from ACID.
Being that it is in-memory, it does not provide durability.
`
