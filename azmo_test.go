package azmo

import (
	"bytes"
	"context"
	"fmt"
	"net"
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

/*
type fmtEncoder struct {
	w io.Writer
}

func (e fmtEncoder) Encode(ev *pb.Event) error {
	_, err := fmt.Fprintln(e.w, ev)
	return err
}
*/

type testEncoder struct {
	typ     pb.Event_Type
	key     []byte
	content []byte
	created int64
	current int64
}

func (e testEncoder) Encode(ev *pb.Event) error {
	if e.typ != ev.Type {
		return fmt.Errorf("expected event type %v, have %v", e.typ, ev.Type)
	}
	if bytes.Compare(e.key, ev.Key) != 0 {
		return fmt.Errorf("expected event key %v, have %v", e.key, ev.Key)
	}
	if bytes.Compare(e.content, ev.Content) != 0 {
		return fmt.Errorf("expected event content %v, have %v", e.content, ev.Content)
	}
	if e.created != ev.Created {
		return fmt.Errorf("expected created rev %d, have %d", e.created, ev.Created)
	}
	if e.current != ev.Current {
		return fmt.Errorf("expected current rev %d, have %d", e.current, ev.Current)
	}
	return nil
}

func TestPutGet(t *testing.T) {
	c, err := Dial(defaultTestServerAddress, 0)
	if err != nil {
		t.Fatalf("cannot dial server: %v", err)
	}
	defer c.Close()

	putTest := testEncoder{typ: pb.Put, key: []byte("k1"), created: 1, current: 1}
	err = c.Put(context.TODO(), putTest, &pb.PutRequest{
		Key:   []byte("k1"),
		Value: []byte("v1"),
	})
	if err != nil {
		t.Fatalf("put k1: %v", err)
	}

	getTest := testEncoder{typ: pb.Get, key: []byte("k1"), content: []byte("v1"), created: 1, current: 1}
	err = c.Get(context.TODO(), getTest, &pb.GetRequest{
		Key:       []byte("k1"),
		Revision:  0,
		MustEqual: false,
	})
	if err != nil {
		t.Fatalf("get k1: %v", err)
	}
}
