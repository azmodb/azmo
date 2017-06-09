package protobuf

import (
	"encoding/binary"
	"io"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

var bufPool = &sync.Pool{New: func() interface{} { return proto.NewBuffer(nil) }}

func nextBuffer(buf []byte) *proto.Buffer {
	b := bufPool.Get().(*proto.Buffer)
	b.SetBuf(buf)
	return b
}

func putBuffer(b *proto.Buffer) {
	if b == nil {
		return
	}
	b.SetBuf(nil)
	bufPool.Put(b)
}

type marshaler interface {
	MarshalTo([]byte) (int, error)
	Size() int
}

func grow(buf []byte, size int) []byte {
	if cap(buf) < size {
		tmp := make([]byte, size)
		copy(tmp, buf)
		buf = tmp
	}
	return buf[:size]
}

const maxInt = uint64(^uint(0) >> 1)

func Marshal(buf []byte, max int, m proto.Message) ([]byte, error) {
	if t, ok := m.(marshaler); ok {
		size := t.Size()
		if max > 0 && size > max {
			return buf, errors.New("protobuf: message too large")
		}
		buf = grow(buf, binary.MaxVarintLen64+size)
		n := binary.PutUvarint(buf, uint64(size))
		done, err := t.MarshalTo(buf[n : n+size])
		n += done
		if err != nil {
			return buf[:n], errors.Wrap(err, "protobuf marshal")
		}
		if done != size {
			return buf[:n], errors.New("protobuf: short marshal")
		}
		return buf[:n], nil
	}

	size := proto.Size(m)
	if max > 0 && size > max {
		return buf, errors.New("protobuf: message too large")
	}
	buf = grow(buf, binary.MaxVarintLen64+size)
	n := binary.PutUvarint(buf, uint64(size))
	b := nextBuffer(buf[n:n : n+size])
	if err := b.Marshal(m); err != nil {
		putBuffer(b)
		return buf, errors.Wrap(err, "protobuf marshal")
	}
	putBuffer(b)
	return buf[:n+size], nil
}

func Write(w *Writer, buf []byte, m proto.Message) ([]byte, error) {
	var err error
	if buf, err = Marshal(buf, w.max, m); err != nil {
		return buf, err
	}

	var done int
	for n := 0; done < len(buf); done += n {
		if n, err = w.Write(buf[done:]); err != nil {
			break
		}

	}
	if err != nil {
		return buf, err
	}
	if done < len(buf) {
		return buf, errors.New("protobuf: short write")
	}
	return buf, nil
}

func Unmarshal(buf []byte, m proto.Message) error {
	v, done := binary.Uvarint(buf)
	switch {
	case done == 0:
		return errors.New("protobuf: uvarint buffer too small")
	case done < 0:
		return errors.New("protobuf: uvarint 64 bit overflow")
	case v > maxInt:
		return errors.New("protobuf: integer overflow")
	}

	off := done + int(v)
	if len(buf) < off {
		return errors.New("protobuf: buffer too small")
	}
	return unmarshal(buf[done:off], m)
}

func Read(r *Reader, buf []byte, m proto.Message) ([]byte, error) {
	v, err := binary.ReadUvarint(r)
	if err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return buf, err
		}
		return buf, errors.Wrap(err, "protobuf")
	}
	if v > maxInt {
		return buf, errors.New("protobuf: integer overflow")
	}
	size := int(v)
	if r.max > 0 && size > r.max {
		return buf, errors.New("protobuf: message too large")
	}

	buf = grow(buf, size)
	n, err := r.Read(buf)
	if err != nil {
		return buf[:n], err
	}
	if n < size {
		return buf[:n], errors.New("protobuf: short read")
	}
	return buf, unmarshal(buf, m)
}

type unmarshaler interface {
	Unmarshal([]byte) error
}

func unmarshal(data []byte, m proto.Message) (err error) {
	if m == nil {
		return nil
	}

	if t, ok := m.(unmarshaler); ok {
		err = t.Unmarshal(data)
	} else {
		err = proto.Unmarshal(data, m)
	}
	if err != nil {
		err = errors.Wrap(err, "protobuf unmarshal")
	}
	return err
}
