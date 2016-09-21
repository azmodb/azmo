package client

import (
	"errors"
	"io"
	"sync"
	"time"

	pb "github.com/azmodb/azmo/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	eventPool = sync.Pool{New: func() interface{} { return &Event{} }}
	rangePool = sync.Pool{New: func() interface{} { return &Range{} }}
	batchPool = sync.Pool{New: func() interface{} { return &Batch{} }}
)

type Event struct {
	*pb.Event
	err error
}

func (e *Event) Err() error { return e.err }

func (e *Event) Close() {
	if e == nil {
		return
	}

	ev := e
	ev.Event = nil
	ev.err = nil
	eventPool.Put(ev)
	e = nil
}

type Range struct {
	from []byte
	to   []byte
	rev  int64
	vers bool
}

func NewRange(from, to []byte, rev int64, versions bool) *Range {
	r := rangePool.Get().(*Range)
	r.from = from
	r.to = to
	r.rev = rev
	r.vers = versions
	return r
}

func (r *Range) Close() {
	if r == nil {
		return
	}

	ra := r
	ra.from = nil
	ra.to = nil
	ra.rev = 0
	rangePool.Put(ra)
	r = nil
}

type Option func(*DB) error

type DB struct {
	conn *grpc.ClientConn
	c    pb.DBClient
}

func Dial(address string, timeout time.Duration, opts ...Option) (*DB, error) {
	options := []grpc.DialOption{grpc.WithInsecure()}
	if timeout > 0 {
		options = append(options, grpc.WithTimeout(timeout))
	}

	conn, err := grpc.Dial(address, options...)
	if err != nil {
		return nil, err
	}
	return NewDB(conn), nil
}

func NewDB(conn *grpc.ClientConn) *DB {
	return &DB{c: pb.NewDBClient(conn), conn: conn}
}

func (db *DB) Close() error {
	if db == nil || db.conn == nil || db.c == nil {
		return errors.New("database is shut down")
	}

	err := db.conn.Close()
	db.conn = nil
	db.c = nil
	return err
}

func (db *DB) Apply(ctx context.Context, b *Batch) (<-chan *Event, error) {
	stream, err := db.c.Batch(ctx, b.req)
	if err != nil {
		return nil, err
	}

	ch := make(chan *Event, b.Len())
	go func(ch chan<- *Event) {
		defer close(ch)

		for {
			resp, err := stream.Recv()
			if err != nil && err == io.EOF {
				break
			}

			ev := eventPool.Get().(*Event)
			ev.Event = resp
			ev.err = err
			ch <- ev
			if err != nil {
				break
			}
		}
	}(ch)

	return ch, nil
}

func (db *DB) Watch(ctx context.Context, key []byte) (<-chan *Event, error) {
	return nil, nil
}

func (db *DB) Get(ctx context.Context, r *Range) (*Event, error) {
	resp, err := db.c.Get(ctx, &pb.GetRequest{
		Key:      r.from,
		Rev:      r.rev,
		Versions: r.vers,
	})

	ev := eventPool.Get().(*Event)
	ev.Event = resp
	ev.err = err
	return ev, err
}

func (db *DB) Range(ctx context.Context, r *Range) (<-chan *Event, error) {
	stream, err := db.c.Range(ctx, &pb.RangeRequest{
		From:     r.from,
		To:       r.to,
		Rev:      r.rev,
		Versions: r.vers,
	})
	if err != nil {
		return nil, err
	}

	ch := make(chan *Event, 32) // TODO: find ch capacity
	for {
		resp, err := stream.Recv()
		if err != nil && err == io.EOF {
			break
		}

		ev := eventPool.Get().(*Event)
		ev.Event = resp
		ev.err = err
		ch <- ev
		if err != nil {
			break
		}
	}
	return ch, nil
}

type Batch struct {
	req *pb.BatchRequest
}

func NewBatch() *Batch {
	b := batchPool.Get().(*Batch)
	if b.req == nil {
		b.req = &pb.BatchRequest{}
	}
	return b
}

func (b *Batch) Close() {
	if b == nil || b.req == nil {
		return
	}

	if b.req.Args != nil {
		b.req.Args = b.req.Args[:0]
	}
	batchPool.Put(b)
}

func (b *Batch) Increment(key []byte, value int64, prev bool) {
	b.req.Args = append(b.req.Args, &pb.Argument{
		Command: &pb.Argument_Increment{
			Increment: &pb.NumericRequest{
				Key:   key,
				Value: value,
				Prev:  prev,
			},
		},
	})
}

func (b *Batch) Decrement(key []byte, value int64, prev bool) {
	b.req.Args = append(b.req.Args, &pb.Argument{
		Command: &pb.Argument_Decrement{
			Decrement: &pb.NumericRequest{
				Key:   key,
				Value: value,
				Prev:  prev,
			},
		},
	})
}

func (b *Batch) Insert(key []byte, value []byte, prev bool) {
	b.req.Args = append(b.req.Args, &pb.Argument{
		Command: &pb.Argument_Put{
			Put: &pb.PutRequest{
				Key:       key,
				Value:     value,
				Prev:      prev,
				Tombstone: false,
			},
		},
	})
}

func (b *Batch) Put(key []byte, value []byte, prev bool) {
	b.req.Args = append(b.req.Args, &pb.Argument{
		Command: &pb.Argument_Put{
			Put: &pb.PutRequest{
				Key:       key,
				Value:     value,
				Prev:      prev,
				Tombstone: true,
			},
		},
	})
}

func (b *Batch) Delete(key []byte, prev bool) {
	b.req.Args = append(b.req.Args, &pb.Argument{
		Command: &pb.Argument_Delete{
			Delete: &pb.DeleteRequest{
				Key:  key,
				Prev: prev,
			},
		},
	})
}

func (b *Batch) Len() int { return len(b.req.Args) }
