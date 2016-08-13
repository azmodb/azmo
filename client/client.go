package client

import (
	"errors"
	"io"
	"time"

	"github.com/azmodb/azmo/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type DB struct {
	conn *grpc.ClientConn
	c    pb.DBClient
}

type Option func(*DB) error

func Dial(address string, timeout time.Duration, options ...Option) (*DB, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if timeout > 0 {
		opts = append(opts, grpc.WithTimeout(timeout))
	}

	conn, err := grpc.Dial(address, opts...)
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

type Pair struct {
	m   interface{}
	key []byte
}

func (p Pair) Value() []byte {
	if p.key == nil || p.m == nil {
		return nil
	}

	if v, ok := p.m.(*pb.RangeResponse); ok {
		return v.Value
	}
	if v, ok := p.m.(*pb.GetResponse); ok {
		return v.Value
	}
	panic("pair message type mismatch")
}

func (p Pair) Revisions() []int64 {
	if p.key == nil || p.m == nil {
		return nil
	}

	if v, ok := p.m.(*pb.RangeResponse); ok {
		return v.Revs
	}
	if v, ok := p.m.(*pb.GetResponse); ok {
		return v.Revs
	}
	panic("pair message type mismatch")
}

func (p Pair) Revision() int64 {
	if p.key == nil || p.m == nil {
		return 0
	}

	if v, ok := p.m.(*pb.RangeResponse); ok {
		return v.Rev
	}
	if v, ok := p.m.(*pb.GetResponse); ok {
		return v.Rev
	}
	panic("pair message type mismatch")
}

func (p Pair) Key() []byte { return p.key }

func (db *DB) Get(ctx context.Context, key []byte, rev int64) (Pair, error) {
	resp, err := db.c.Get(ctx, &pb.GetRequest{Key: key, Rev: rev})
	if err != nil {
		return Pair{}, err
	}
	if resp.Value == nil {
		return Pair{}, errors.New("key or revision not found")
	}
	return Pair{key: key, m: resp}, nil
}

type Range [2][]byte

func NewRange(from, to []byte) Range { return [2][]byte{from, to} }

func (db *DB) Range(ctx context.Context, r Range, rev int64) ([]Pair, error) {
	stream, err := db.c.Range(ctx, &pb.RangeRequest{
		From: r[0],
		To:   r[1],
		Rev:  rev,
	})
	if err != nil {
		return nil, err
	}

	pairs := make([]Pair, 0, 16) // TODO
	for {
		resp, err := stream.Recv()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, Pair{key: resp.Key, m: resp})
	}
	return pairs, nil
}

type Result struct {
	m   interface{}
	key []byte
}

func (r Result) Key() []byte { return r.key }

func (r Result) Revisions() []int64 {
	if m, ok := r.m.(*pb.DeleteResponse); ok {
		return m.Revs
	}
	if m, ok := r.m.(*pb.PutResponse); ok {
		return m.Revs
	}
	panic("txn result message type mismatch")
}

func (r Result) Revision() int64 {
	if m, ok := r.m.(*pb.DeleteResponse); ok {
		return m.Rev
	}
	if m, ok := r.m.(*pb.PutResponse); ok {
		return m.Rev
	}
	panic("txn result message type mismatch")
}

func (db *DB) Insert(ctx context.Context, key, value []byte) (Result, error) {
	resp, err := db.c.Put(ctx, &pb.PutRequest{Key: key, Value: value, Tombstone: true})
	if err != nil {
		return Result{}, err
	}
	return Result{key: key, m: resp}, nil
}

func (db *DB) Put(ctx context.Context, key, value []byte) (Result, error) {
	resp, err := db.c.Put(ctx, &pb.PutRequest{Key: key, Value: value})
	if err != nil {
		return Result{}, err
	}
	return Result{key: key, m: resp}, nil
}

func (db *DB) Delete(ctx context.Context, key []byte) (Result, error) {
	resp, err := db.c.Delete(ctx, &pb.DeleteRequest{Key: key})
	if err != nil {
		return Result{}, err
	}
	return Result{key: key, m: resp}, nil
}

func (db *DB) Txn(ctx context.Context, batch *TxnBatch) ([]TxnResult, error) {
	resp, err := db.c.Txn(ctx, batch.txnRequest())
	if err != nil {
		return nil, err
	}

	results := make([]TxnResult, 0, len(resp.Responses))
	for _, r := range resp.Responses {
		req := batch.request(r.Num)
		if req == nil {
			return nil, errors.New("txnid not found")
		}

		results = append(results, TxnResult{
			key: req.Key,
			m:   r,
		})
	}
	return results, nil
}
