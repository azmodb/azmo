package protobuf

import (
	"bytes"
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/azmodb/protobuf/custompb"
	"github.com/azmodb/protobuf/googlepb"
	"github.com/golang/protobuf/proto"
)

var (
	stringBlob = "Build a better mousetrap, and nature will build a better mouse."
	bytesBlob  = []byte(stringBlob)
)

var compatTestMessages = []proto.Message{
	&googlepb.Message{stringBlob, math.MaxUint64, math.MaxUint32, bytesBlob},
	&googlepb.Message{"test-string", 42, 42, []byte("test-bytes")},
	&googlepb.Message{"", 0, 0, nil},
	&googlepb.Message{},

	&custompb.Message{stringBlob, math.MaxUint64, math.MaxUint32, bytesBlob},
	&custompb.Message{"test-string", 42, 42, []byte("test-bytes")},
	&custompb.Message{"", 0, 0, nil},
	&custompb.Message{},
}

func newMessageFromType(in proto.Message) proto.Message {
	if _, ok := in.(*custompb.Message); ok {
		return &custompb.Message{}
	}
	return &googlepb.Message{}
}

func deepEqual(in, out proto.Message) error {
	if t, ok := in.(*custompb.Message); ok {
		if !reflect.DeepEqual(t, out.(*custompb.Message)) {
			return errors.New("custom message differ")
		}
		return nil
	}
	if !reflect.DeepEqual(in.(*googlepb.Message), out.(*googlepb.Message)) {
		return errors.New("google message differ")
	}
	return nil
}

func TestBasicMarshalUnmarshal(t *testing.T) {
	var buf []byte
	var err error

	for i, in := range compatTestMessages {
		if buf, err = Marshal(buf, 0, in); err != nil {
			t.Fatalf("codec test #%d: marshal failed: %v", i, err)
		}

		out := newMessageFromType(in)
		if err = Unmarshal(buf, out); err != nil {
			t.Fatalf("codec test #%d: unmarshal failed: %v", i, err)
		}
		if err = deepEqual(in, out); err != nil {
			t.Fatalf("marshal/unmarshal #%d: %v", i, err)
		}

		out.Reset()
		r := NewReader(bytes.NewBuffer(buf), 0)
		if _, err = Read(r, nil, out); err != nil {
			t.Fatalf("codec test #%d: read failed: %v", i, err)
		}
		if err = deepEqual(in, out); err != nil {
			t.Fatalf("marshal/unmarshal #%d: %v", i, err)
		}
	}
}

func TestBasicWriteRead(t *testing.T) {
	var buf []byte
	var err error

	for i, in := range compatTestMessages {
		b := bytes.NewBuffer(nil)
		w := NewWriter(b, 0)
		if buf, err = Write(w, buf, in); err != nil {
			t.Fatalf("codec test #%d: write failed: %v", i, err)
		}
		if err = w.Flush(); err != nil {
			t.Fatalf("codec test #%d: flush failed: %v", i, err)
		}

		out := newMessageFromType(in)
		r := NewReader(b, 0)
		if _, err = Read(r, nil, out); err != nil {
			t.Fatalf("codec test #%d: read failed: %v", i, err)
		}
		if err = deepEqual(in, out); err != nil {
			t.Fatalf("write/read #%d: %v", i, err)
		}
	}
}

func TestBasicEncodeDecode(t *testing.T) {
	var err error

	for i, in := range compatTestMessages {
		b := bytes.NewBuffer(nil)
		enc := NewEncoder(b, 0)
		if err = enc.Encode(in); err != nil {
			t.Fatalf("codec test #%d: encode failed: %v", i, err)
		}
		if err = enc.Flush(); err != nil {
			t.Fatalf("codec test #%d: flush failed: %v", i, err)
		}

		out := newMessageFromType(in)
		dec := NewDecoder(b, 0)
		if err = dec.Decode(out); err != nil {
			t.Fatalf("codec test #%d: decode failed: %v", i, err)
		}
		if err = deepEqual(in, out); err != nil {
			t.Fatalf("encode/decode #%d: %v", i, err)
		}
	}
}
