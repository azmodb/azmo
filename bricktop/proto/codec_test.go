package proto

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	var tests = []struct {
		in, out Fcall
	}{
		{&Tattach{Fid: 2, Afid: 1, Uname: "glenda", Aname: "/usr/glenda", Uid: 42}, &Tattach{}},

		{&Tversion{Version: "9P2000.L", Msize: 8192}, &Tversion{}},
		{&Rversion{Version: "9P2000.L", Msize: 8192}, &Rversion{}},
	}

	buf := &bytes.Buffer{}
	enc := NewEncoder(buf, 8192)
	dec := NewDecoder(buf, 8192)
	h := Header{}
	for i, test := range tests {
		tag := uint16(i) + 1

		if err := enc.Encode(tag, test.in); err != nil {
			t.Fatalf("encode #%d: unexpected encode error: %v", i, err)
		}

		h = Header{}
		if err := dec.DecodeHeader(&h); err != nil {
			t.Fatalf("decode header #%d: unexpected error: %v", i, err)
		}
		if err := dec.Decode(&h, test.out); err != nil {
			t.Fatalf("decode #%d: unexpected decode error: %v", i, err)
		}

		if !reflect.DeepEqual(test.in, test.out) {
			t.Fatalf("encode/decode #%d: result differ\nin  %+v\nout %+v",
				test.in, test.out)
		}
	}
}
