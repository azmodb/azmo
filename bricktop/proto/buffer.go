package proto

import (
	"errors"
	"io"
)

// Buffer is a buffer manager for marshaling and unmarshaling 9P protocol
// messages. It may be reused between invocations to reduce memory usage.
//
// If an error occurs encoding to or decoding from a Buffer, no more data
// will be accepted and all subsequent calls will return the error.
type Buffer struct {
	data    []byte
	err     error
	maxsize int
}

// memory to hold first slice; helps small buffers avoid allocation
const minBufLen = 256

func NewBuffer(data []byte, minsize int, maxsize int) *Buffer {
	if len(data) < 256 {
		data = make([]byte, 0, minBufLen)
	}
	return &Buffer{data: data, maxsize: maxsize}
}

// Reset resets the buffer to be empty, but it retains the underlying
// storage for use by future writes.
func (b *Buffer) Reset() {
	b.data = b.data[0:0]
	b.err = nil
}

// Len returns the number of bytes of the unread portion of the buffer.
func (b *Buffer) Len() int { return len(b.data) }

// Cap returns the capacity of the buffer's underlying byte slice, that
// is, the total space allocated for the buffer's data.
func (b *Buffer) Cap() int { return cap(b.data) }

// Err returns the error, if any, that was encountered during encoding
// or decoding.
func (b *Buffer) Err() error { return b.err }

// Bytes returns a slice of length b.Len() holding the unread portion of
// the buffer. The slice is valid for use only until the next buffer
// modification
func (b *Buffer) Bytes() []byte { return b.data }

// EncodeUint64 writes a little-endian encoded 64-bit integer to the
// Buffer.
func (b *Buffer) EncodeUint64(v uint64) error {
	if b.err != nil {
		return b.err
	}

	n, err := b.grow(8)
	if err != nil {
		b.err = err
		return err
	}

	b.data[n] = byte(v)
	b.data[n+1] = byte(v >> 8)
	b.data[n+2] = byte(v >> 16)
	b.data[n+3] = byte(v >> 24)
	b.data[n+4] = byte(v >> 32)
	b.data[n+5] = byte(v >> 40)
	b.data[n+6] = byte(v >> 48)
	b.data[n+7] = byte(v >> 56)
	return nil
}

// EncodeUint32 writes a little-endian encoded 32-bit integer to the
// Buffer.
func (b *Buffer) EncodeUint32(v uint32) error {
	if b.err != nil {
		return b.err
	}

	n, err := b.grow(4)
	if err != nil {
		b.err = err
		return err
	}

	b.data[n] = byte(v)
	b.data[n+1] = byte(v >> 8)
	b.data[n+2] = byte(v >> 16)
	b.data[n+3] = byte(v >> 24)
	return nil
}

// EncodeUint16 writes a little-endian encoded 16-bit integer to the
// Buffer.
func (b *Buffer) EncodeUint16(v uint16) error {
	if b.err != nil {
		return b.err
	}

	n, err := b.grow(2)
	if err != nil {
		b.err = err
		return err
	}

	b.data[n] = byte(v)
	b.data[n+1] = byte(v >> 8)
	return nil
}

// EncodeUint8 writes a little-endian encoded 8-bit integer to the
// Buffer.
func (b *Buffer) EncodeUint8(v uint8) error {
	if b.err != nil {
		return b.err
	}

	n, err := b.grow(1)
	if err != nil {
		b.err = err
		return err
	}

	b.data[n] = v
	return nil
}

// EncodeString writes a 16-bit count-delimited string slice to the
// Buffer.
func (b *Buffer) EncodeString(v string) error {
	if b.err != nil {
		return b.err
	}

	size := len(v)
	if size > 1<<16-1 {
		b.err = errors.New("name too large")
		return b.err
	}

	n, err := b.grow(2 + size)
	if err != nil {
		b.err = err
		return err
	}

	b.data[n] = byte(size)
	b.data[n+1] = byte(size >> 8)
	copy(b.data[n+2:], v)
	return nil
}

// EncodeBytes writes a 32-bit count-delimited byte slice to the Buffer.
func (b *Buffer) EncodeBytes(v []byte) error {
	if b.err != nil {
		return b.err
	}

	size := len(v)
	if size > 1<<32-1 {
		b.err = errors.New("data slice too large")
		return b.err
	}

	n, err := b.grow(4 + size)
	if err != nil {
		b.err = err
		return err
	}

	b.data[n] = byte(size)
	b.data[n+1] = byte(size >> 8)
	b.data[n+2] = byte(size >> 16)
	b.data[n+3] = byte(size >> 24)
	copy(b.data[n+4:], v)
	return nil
}

func (b *Buffer) grow(size int) (int, error) {
	n := len(b.data)
	if n+size > b.maxsize {
		return 0, errors.New("raw data slice too large")
	}

	if n+size > cap(b.data) {
		m := size + 2*cap(b.data)
		data := make([]byte, n, m)
		copy(data, b.data)
		b.data = data
	}
	b.data = b.data[0 : n+size]
	return n, nil
}

// DecodeUint64 reads a little-endian encoded 64-bit integer from the
// Buffer.
func (b *Buffer) DecodeUint64(v *uint64) error {
	if b.err != nil {
		return b.err
	}
	if len(b.data) < 8 {
		b.err = io.ErrUnexpectedEOF
		return b.err
	}

	*v = uint64(b.data[0]) |
		uint64(b.data[1])<<8 |
		uint64(b.data[2])<<16 |
		uint64(b.data[3])<<24 |
		uint64(b.data[4])<<32 |
		uint64(b.data[5])<<40 |
		uint64(b.data[6])<<48 |
		uint64(b.data[7])<<56
	b.data = b.data[8:]
	return nil
}

// DecodeUint32 reads a little-endian encoded 32-bit integer from the
// Buffer.
func (b *Buffer) DecodeUint32(v *uint32) error {
	if b.err != nil {
		return b.err
	}
	if len(b.data) < 4 {
		b.err = io.ErrUnexpectedEOF
		return b.err
	}

	*v = uint32(b.data[0]) |
		uint32(b.data[1])<<8 |
		uint32(b.data[2])<<16 |
		uint32(b.data[3])<<24
	b.data = b.data[4:]
	return nil
}

// DecodeUint16 reads a little-endian encoded 16-bit integer from the
// Buffer.
func (b *Buffer) DecodeUint16(v *uint16) error {
	if b.err != nil {
		return b.err
	}
	if len(b.data) < 2 {
		b.err = io.ErrUnexpectedEOF
		return b.err
	}

	*v = uint16(b.data[0]) | uint16(b.data[1])<<8
	b.data = b.data[2:]
	return nil
}

// DecodeUint8 reads a little-endian encoded 8-bit integer from the
// Buffer.
func (b *Buffer) DecodeUint8(v *uint8) error {
	if b.err != nil {
		return b.err
	}
	if len(b.data) < 1 {
		b.err = io.ErrUnexpectedEOF
		return b.err
	}

	*v = b.data[0]
	b.data = b.data[1:]
	return nil
}

// DecodeString reads a 16-bit count-delimited encoded string from the
// Buffer.
func (b *Buffer) DecodeString(v *string) error {
	if b.err != nil {
		return b.err
	}
	if len(b.data) < 2 {
		b.err = io.ErrUnexpectedEOF
		return b.err
	}

	size := uint16(b.data[0]) | uint16(b.data[1])<<8
	if size > 1<<16-1 {
		return errors.New("name too large")
	}

	*v = string(b.data[2 : 2+size])
	b.data = b.data[2+size:]
	return nil
}

// DecodeBytes reads a 32-bit count-delimited encoded bytes slice from
// the Buffer.
func (b *Buffer) DecodeBytes(v *[]byte) error {
	if b.err != nil {
		return b.err
	}
	if len(b.data) < 4 {
		b.err = io.ErrUnexpectedEOF
		return b.err
	}

	size := uint32(b.data[0]) |
		uint32(b.data[1])<<8 |
		uint32(b.data[2])<<16 |
		uint32(b.data[3])<<24
	if size > 1<<32-1 {
		b.err = errors.New("data slice too large")
		return b.err
	}

	if cap(*v) < int(size) {
		*v = make([]byte, size)
	}
	*v = (*v)[:size]
	copy(*v, b.data[4:4+size])
	b.data = b.data[4+size:]
	return nil

}
