package proto

import (
	"bufio"
	"errors"
	"io"
)

type Decoder struct {
	r       io.Reader
	buf     *Buffer
	maxsize int
}

func NewDecoder(r io.Reader, maxsize uint32) *Decoder {
	if maxsize > MaxMessageSize {
		maxsize = MaxMessageSize
	}
	if maxsize < MaxHeaderSize {
		maxsize = MaxMessageSize
	}

	return &Decoder{
		buf:     NewBuffer(nil, 256, int(maxsize)),
		r:       bufio.NewReader(r),
		maxsize: int(maxsize),
	}
}

func (d *Decoder) DecodeHeader(h *Header) error {
	d.buf.data = d.buf.data[0:7]
	_, err := io.ReadFull(d.r, d.buf.data[:7])
	if err != nil {
		return err
	}

	d.buf.DecodeUint32(&h.msize)
	d.buf.DecodeUint8(&h.Type)
	d.buf.DecodeUint16(&h.Tag)
	if err = d.buf.Err(); err != nil {
		return err
	}
	if h.msize < 7 {
		return errors.New("fcall too small")
	}
	if h.msize > MaxMessageSize {
		return errors.New("fcall too large")
	}
	return err
}

func (d *Decoder) Decode(h *Header, f Fcall) error {
	if h.msize < 7 {
		return errors.New("fcall too small")
	}
	if h.msize > MaxMessageSize {
		return errors.New("fcall too large")
	}

	d.buf.data = d.buf.data[0:0]
	msize := int(h.msize) - 7
	_, err := d.buf.grow(msize)
	if err != nil {
		return err
	}
	_, err = io.ReadFull(d.r, d.buf.data[:msize])
	if err != nil {
		return err
	}

	if f == nil {
		return nil
	}
	return unmarshal(d.buf, h.Type, h.Tag, f)
}

type Encoder struct {
	w       *bufio.Writer
	buf     *Buffer
	msize   [4]byte
	maxsize int
}

func NewEncoder(w io.Writer, maxsize uint32) *Encoder {
	if maxsize > MaxMessageSize {
		maxsize = MaxMessageSize
	}
	if maxsize < MaxHeaderSize {
		maxsize = MaxMessageSize
	}

	return &Encoder{
		buf:     NewBuffer(nil, 256, int(maxsize)),
		w:       bufio.NewWriter(w),
		maxsize: int(maxsize),
	}
}

func (e *Encoder) Encode(tag uint16, f Fcall) error {
	e.buf.data = e.buf.data[0:0]
	err := marshal(e.buf, tag, f)
	if err != nil {
		return err
	}

	msize := 4 + uint32(e.buf.Len())
	e.msize[0] = byte(msize)
	e.msize[1] = byte(msize >> 8)
	e.msize[2] = byte(msize >> 16)
	e.msize[3] = byte(msize >> 24)
	_, err = e.w.Write(e.msize[:])
	if err != nil {
		return err
	}
	_, err = e.w.Write(e.buf.Bytes())
	if err != nil {
		return err
	}
	return e.w.Flush()
}
