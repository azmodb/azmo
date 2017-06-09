package azmo

import (
	"bytes"
	"fmt"
	"net"
	"testing"

	pb "github.com/azmodb/azmo/azmopb"
	"github.com/azmodb/db"
	"golang.org/x/net/context"
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

func verify(want, ev *pb.Event) error {
	if want.Type != ev.Type {
		return fmt.Errorf("expected event type %v, have %v", want.Type, ev.Type)
	}
	if bytes.Compare(want.Key, ev.Key) != 0 {
		return fmt.Errorf("expected event key %v, have %v", want.Key, ev.Key)
	}
	if bytes.Compare(want.Content, ev.Content) != 0 {
		return fmt.Errorf("expected event content %v, have %v", want.Content, ev.Content)
	}
	if want.Created != ev.Created {
		return fmt.Errorf("expected created rev %d, have %d", want.Created, ev.Created)
	}
	if want.Current != ev.Current {
		return fmt.Errorf("expected current rev %d, have %d", want.Current, ev.Current)
	}
	return nil
}

func TestPutGet(t *testing.T) {
	c, err := Dial(defaultTestServerAddress, 0)
	if err != nil {
		t.Fatalf("cannot dial server: %v", err)
	}
	defer c.Close()

	key, val := []byte("put_get_test_1"), []byte("v1")
	want := &pb.Event{Type: pb.Put, Key: key, Created: 1, Current: 1}
	ev, err := c.Put(context.TODO(), &pb.PutRequest{
		Key:   key,
		Value: val,
	})
	if err != nil {
		t.Fatalf("put put_get_test_1: %v", err)
	}
	if err := verify(want, ev.Event); err != nil {
		t.Fatalf("put/get: expected event differ:\n%v\n%v", want, ev.Event)
	}

	want = &pb.Event{Type: pb.Get, Key: key, Content: val, Created: 1, Current: 1}
	ev, err = c.Get(context.TODO(), &pb.GetRequest{
		Key:       []byte("put_get_test_1"),
		Revision:  0,
		MustEqual: false,
	})
	if err != nil {
		t.Fatalf("get k1: %v", err)
	}
	if err := verify(want, ev.Event); err != nil {
		t.Fatalf("put/get: expected event differ:\n%v\n%v", want, ev.Event)
	}
}

/*
func TestRange(t *testing.T) {
	c, err := Dial(defaultTestServerAddress, 0)
	if err != nil {
		t.Fatalf("cannot dial server: %v", err)
	}
	defer c.Close()

	noop := noopEncoder{}
	count := 10
	for i := 0; i < count; i++ {
		key := []byte(fmt.Sprintf("range_test_key:%.4d", i))
		err = c.Put(context.TODO(), noop, &pb.PutRequest{
			Key:   key,
			Value: []byte("v1"),
		})
		if err != nil {
			t.Fatalf("put %s: %v", key, err)
		}
	}

	enc := fmtEncoder{os.Stdout}
	from := []byte("range_test_key:0000")
	to := []byte(fmt.Sprintf("range_test_key:%.4d", count))
	err = c.Range(context.TODO(), enc, &pb.RangeRequest{
		From:     from,
		To:       to,
		Revision: 0,
		Limit:    0,
	})
	if err != nil {
		t.Fatalf("range: %v", err)
	}
}
*/
