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

type Option func()

func Dial(address string, timeout time.Duration, options ...Option) (*DB, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &DB{c: pb.NewDBClient(conn), conn: conn}, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
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

type TxnResult struct {
	resps []*pb.Response
	keys  [][]byte
}

func (r TxnResult) Len() int { return len(r.resps) }

func (r *TxnResult) Next() (key []byte, num int32, rev int64, ok bool) {
	if r == nil || len(r.resps) == 0 || len(r.keys) == 0 {
		return key, 0, 0, false
	}

	resp := r.resps[0]
	key = r.keys[0]
	r.resps = r.resps[1:]
	r.keys = r.keys[1:]

	num = resp.Num
	rev = resp.Rev
	return key, num, rev, true
}

func (db *DB) Apply(ctx context.Context, b *TxnBatch) (*TxnResult, error) {
	resp, err := db.c.Txn(ctx, b.txnRequest())
	if err != nil {
		return nil, err
	}
	if len(resp.Responses) != b.Len() {
		return nil, errors.New("result length mismatch")
	}

	return &TxnResult{resps: resp.Responses, keys: b.keys()}, nil
}
