package azmo

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"

	pb "github.com/azmodb/azmo/azmopb"
	"github.com/azmodb/db"
)

const defaultTestServerAddress = "localhost:7979"

var listener net.Listener

func newServer(addr string) (listener net.Listener, err error) {
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := Listen(db.New(), listener, "", ""); err != nil {
			panic(fmt.Sprintf("azmo server failed: %v", err))
		}
	}()
	return listener, err
}

func init() {
	l, err := newServer(defaultTestServerAddress)
	if err != nil {
		panic(fmt.Sprintf("azmo server failed: %v", err))
	}
	listener = l
}

func TestBasicServerClient(t *testing.T) {
	c, err := Dial(defaultTestServerAddress, 0)
	if err != nil {
		t.Fatalf("cannot dial server: %v", err)
	}
	defer c.Close()

	err = c.Put(context.TODO(), os.Stdout, &pb.PutRequest{
		Key:   []byte("k1"),
		Value: []byte("v1"),
	})
	if err != nil {
		t.Fatalf("put k1: %v", err)
	}

	err = c.Get(context.TODO(), os.Stdout, &pb.GetRequest{
		Key:       []byte("k1"),
		Revision:  0,
		MustEqual: false,
	})
	if err != nil {
		t.Fatalf("get k1: %v", err)
	}
}
