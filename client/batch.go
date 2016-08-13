package client

import (
	"sort"

	"github.com/azmodb/azmo/pb"
)

type requests []*pb.Request

func (r *requests) Put(num int32, key, value []byte, tombstone bool) {
	*r = append(*r, &pb.Request{
		Type:      pb.Request_PutRequest,
		Num:       num,
		Key:       key,
		Value:     value,
		Tombstone: tombstone,
	})
}

func (r *requests) Delete(num int32, key []byte) {
	*r = append(*r, &pb.Request{
		Type: pb.Request_DeleteRequest,
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

func (b *TxnBatch) Put(key, value []byte, tombstone bool) {
	b.sorted = false
	b.num++
	b.reqs.Put(b.num, key, value, tombstone)
}

func (b *TxnBatch) Delete(key []byte) {
	b.sorted = false
	b.num++
	b.reqs.Delete(b.num, key)
}

func (b *TxnBatch) Len() int { return len(b.reqs) }

func (b *TxnBatch) Sort() {
	if !b.sorted {
		sort.Sort(b.reqs)
	}
}

func (b *TxnBatch) txnRequest() *pb.TxnRequest {
	if b == nil || b.reqs == nil {
		return nil
	}

	b.Sort()
	return &pb.TxnRequest{Requests: b.reqs}
}

func (b *TxnBatch) request(num int32) *pb.Request {
	if b == nil || b.reqs == nil {
		return nil
	}

	b.Sort()
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

func (b *TxnBatch) keys() [][]byte {
	keys := make([][]byte, 0, len(b.reqs))
	for _, req := range b.reqs {
		keys = append(keys, req.Key)
	}
	return keys
}
