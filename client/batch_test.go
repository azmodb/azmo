package client

import (
	"bytes"
	"testing"
)

func TestTxnBatchRequest(t *testing.T) {
	b := NewTxnBatch()
	b.Delete([]byte("k1"))
	b.Delete([]byte("k2"))
	b.Delete([]byte("k3"))
	b.Delete([]byte("k4"))
	b.Delete([]byte("k5"))

	for _, test := range []struct {
		key []byte
		num int32
	}{
		{key: nil, num: 0},

		{key: []byte("k1"), num: 1},
		{key: []byte("k2"), num: 2},
		{key: []byte("k3"), num: 3},
		{key: []byte("k4"), num: 4},
		{key: []byte("k5"), num: 5},

		{key: nil, num: 6},
	} {
		req := b.request(test.num)
		if req == nil {
			if test.key != nil {
				t.Fatalf("batch request: expected <nil> key, got %q", req.Key)
			}
			continue
		}
		if bytes.Compare(req.Key, test.key) != 0 {
			t.Fatalf("batch request: expected key %q, got %q", test.key, req.Key)
		}
		if req.Num != test.num {
			t.Fatalf("batch request: expected num %d, got %d", test.num, req.Num)
		}
	}
}
