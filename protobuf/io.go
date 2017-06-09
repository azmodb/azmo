package protobuf

import (
	"bufio"
	"io"

	"github.com/golang/protobuf/proto"
)

const maxKeepAroundBufferLen = 2 * 8192

type Encoder struct {
	w   *Writer
	buf []byte
}

func NewEncoder(w io.Writer, max int) *Encoder {
	return &Encoder{
		buf: make([]byte, 0, 1024),
		w:   NewWriter(w, max),
	}
}

func (e *Encoder) Encode(m proto.Message) (err error) {
	e.buf, err = Write(e.w, e.buf, m)
	if cap(e.buf) > maxKeepAroundBufferLen {
		e.buf = e.buf[0:0:maxKeepAroundBufferLen]
	} else {
		e.buf = e.buf[0:0]
	}
	return err
}

func (e *Encoder) Flush() error { return e.w.Flush() }

type Decoder struct {
	r   *Reader
	buf []byte
}

func NewDecoder(r io.Reader, max int) *Decoder {
	return &Decoder{
		buf: make([]byte, 0, 1024),
		r:   NewReader(r, max),
	}
}

func (d *Decoder) Decode(m proto.Message) (err error) {
	d.buf, err = Read(d.r, d.buf, m)
	if cap(d.buf) > maxKeepAroundBufferLen {
		d.buf = d.buf[0:0:maxKeepAroundBufferLen]
	} else {
		d.buf = d.buf[0:0]
	}
	return err
}

type flusher interface {
	Flush() error
}

type Writer struct {
	w   *bufio.Writer
	err error
	max int
}

func NewWriter(w io.Writer, max int) *Writer {
	bw := &Writer{max: max}
	if t, ok := w.(*bufio.Writer); ok {
		bw.w = t
	} else {
		bw.w = bufio.NewWriter(w)
	}
	return bw
}

func (w *Writer) Write(p []byte) (n int, err error) {
	if w.err != nil {
		err = w.err
		return n, err
	}
	if n, err = w.w.Write(p); err != nil {
		w.err = err
	}
	return n, err
}

func (w *Writer) Flush() (err error) {
	if w.err != nil {
		err = w.err
		return err
	}
	if err = w.w.Flush(); err != nil {
		w.err = err
	}
	return err
}

type byteReader interface {
	io.ByteReader
	io.Reader
}

type Reader struct {
	r   byteReader
	err error
	max int
}

func NewReader(r io.Reader, max int) *Reader {
	br := &Reader{max: max}
	if t, ok := r.(byteReader); ok {
		br.r = t
	} else {
		br.r = bufio.NewReader(r)
	}
	return br
}

func (r *Reader) ReadByte() (b byte, err error) {
	if r.err != nil {
		err = r.err
		return b, err
	}
	if b, err = r.r.ReadByte(); err != nil {
		r.err = err
	}
	return b, err
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.err != nil {
		err = r.err
		return n, err
	}
	if n, err = r.r.Read(p); err != nil {
		r.err = err
	}
	return n, err
}
