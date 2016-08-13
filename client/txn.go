package client

import (
	"sort"

	"github.com/azmodb/azmo/pb"
)

type requests []*pb.GenericRequest

func (r *requests) Put(num int32, key, value []byte, tombstone bool) {
	*r = append(*r, &pb.GenericRequest{
		Type:      pb.GenericRequest_PutRequest,
		Num:       num,
		Key:       key,
		Value:     value,
		Tombstone: tombstone,
	})
}

func (r *requests) Delete(num int32, key []byte) {
	*r = append(*r, &pb.GenericRequest{
		Type: pb.GenericRequest_DeleteRequest,
		Num:  num,
		Key:  key,
	})
}

func (r requests) Less(i, j int) bool { return r[i].Num < r[j].Num }
func (r requests) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r requests) Len() int           { return len(r) }

type TxnBatch struct {
	reqs   requests
	num    int32
	sorted bool
}

func NewTxnBatch() *TxnBatch { return &TxnBatch{reqs: requests{}} }

func (b *TxnBatch) Insert(key, value []byte) {
	b.sorted = false
	b.num++
	b.reqs.Put(b.num, key, value, true)
}

func (b *TxnBatch) Put(key, value []byte) {
	b.sorted = false
	b.num++
	b.reqs.Put(b.num, key, value, false)
}

func (b *TxnBatch) Delete(key []byte) {
	b.sorted = false
	b.num++
	b.reqs.Delete(b.num, key)
}

func (b *TxnBatch) Len() int { return len(b.reqs) }

func (b *TxnBatch) sort() {
	if !b.sorted {
		sort.Sort(b.reqs)
	}
}

func (b *TxnBatch) txnRequest() *pb.TxnRequest {
	if b == nil || b.reqs == nil {
		return nil
	}

	b.sort()
	return &pb.TxnRequest{Requests: b.reqs}
}

func (b *TxnBatch) request(num int32) *pb.GenericRequest {
	if b == nil || b.reqs == nil {
		return nil
	}

	b.sort()
	i := sort.Search(len(b.reqs), func(i int) bool {
		return b.reqs[i].Num >= num
	})
	if i >= len(b.reqs) {
		return nil
	}
	if b.reqs[i].Num == num {
		return b.reqs[i]
	}
	return nil
}

type TxnResult struct {
	m   *pb.GenericResponse
	key []byte
}

func (r TxnResult) Key() []byte { return r.key }

func (r TxnResult) Num() int32 {
	if r.key == nil || r.m == nil {
		return 0
	}
	return r.m.Num
}

func (r TxnResult) Revisions() []int64 {
	if r.key == nil || r.m == nil {
		return nil
	}
	return r.m.Revs
}

func (r TxnResult) Revision() int64 {
	if r.key == nil || r.m == nil {
		return 0
	}
	return r.m.Rev
}
