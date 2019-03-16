package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"github.com/azmodb/bricktop/proto"
)

type Client struct {
	writer sync.Mutex // exclusive writer lock
	enc    *proto.Encoder

	dec *proto.Decoder

	mu       sync.Mutex // protects following
	freetag  map[uint16]struct{}
	tag      uint16
	pending  map[uint16]*Fcall
	closing  bool
	shutdown bool
	c        io.Closer
}

func Dial(ctx context.Context, network, address string) (*Client, error) {
	conn, err := (&net.Dialer{}).DialContext(ctx, network, address)
	if err != nil {
		return nil, err
	}

	c := NewClient(conn)
	if err = c.Version(ctx, proto.MaxMessageSize); err != nil {
		c.Close()
		return nil, err
	}
	return c, nil
}

func NewClient(rwc io.ReadWriteCloser) *Client {
	c := &Client{
		enc: proto.NewEncoder(rwc, proto.MaxMessageSize),
		dec: proto.NewDecoder(rwc, proto.MaxMessageSize),

		freetag: make(map[uint16]struct{}),
		tag:     1,
		pending: make(map[uint16]*Fcall),
		c:       rwc,
	}
	go c.recv()
	return c
}

var errShutdown = errors.New("connection is shut down")

func (c *Client) Close() error {
	c.mu.Lock()
	if c.closing {
		c.mu.Unlock()
		return errShutdown
	}
	c.closing = true
	c.mu.Unlock()
	return c.c.Close()
}

type Fcall struct {
	Args  proto.Fcall
	Reply proto.Fcall
	err   error
	ch    chan<- *Fcall
}

func (f *Fcall) Err() error { return f.err }

func (f *Fcall) done() {
	select {
	case f.ch <- f:
	default:
		log.Println("discarding reply due to insufficient capacity")
	}
}

func (c *Client) nextTag() (uint16, error) {
	for tag, _ := range c.freetag {
		delete(c.freetag, tag)
		return tag, nil
	}

	tag := c.tag
	if tag >= 1<<16-1 {
		return 0, errors.New("out of tags")
	}
	c.tag++
	return tag, nil
}

func (c *Client) Do(ctx context.Context, f *Fcall, ch chan<- *Fcall) {
	f.ch = ch
	select {
	case <-ctx.Done():
		f.err = ctx.Err()
		f.done()
	default:
		c.writer.Lock()
		c.send(f)
		c.writer.Unlock()
		log.Printf("<- %+v", f.Args)
	}
}

func (c *Client) send(f *Fcall) {
	c.mu.Lock()
	if c.shutdown || c.closing {
		c.mu.Unlock()
		f.err = errShutdown
		f.done()
		return
	}
	tag, err := c.nextTag()
	fmt.Println(tag)
	if err != nil {
		c.mu.Unlock()
		f.err = err
		f.done()
		return
	}
	c.pending[tag] = f
	c.mu.Unlock()

	if err := c.enc.Encode(tag, f.Args); err != nil {
		c.mu.Lock()
		f = c.pending[tag]
		delete(c.pending, tag)
		c.freetag[tag] = struct{}{}
		c.mu.Unlock()
		if f != nil {
			f.err = err
			f.done()
		}
	}
}

func (c *Client) recv() {
	var h proto.Header
	var err error
	for err == nil {
		if err = c.dec.DecodeHeader(&h); err != nil {
			break
		}

		c.mu.Lock() // lookup up and delete pending fcall
		tag := h.Tag
		f := c.pending[tag]
		delete(c.pending, tag)
		c.freetag[tag] = struct{}{}
		c.mu.Unlock()

		switch {
		case f == nil:
			err = c.dec.Decode(&h, nil)
			continue
		case h.Type == 7:
			f.Reply = &proto.Rerror{}
			if err = c.dec.Decode(&h, f.Reply); err != nil {
				f.err = err
			}
			f.err = errors.New("TODO: received err code")
			f.done()
		default:
			if err = c.dec.Decode(&h, f.Reply); err != nil {
				f.err = err
			}
			f.done()
		}
	}

	c.writer.Lock()
	c.mu.Lock()
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		if c.closing {
			err = errShutdown
		} else {
			err = io.ErrUnexpectedEOF
		}
	}
	for tag, f := range c.pending {
		f.err = err
		f.done()
		delete(c.pending, tag)
		c.freetag[tag] = struct{}{}
	}
	c.mu.Unlock()
	c.writer.Unlock()
}

const version = "9P2000.L"

func (c *Client) Version(ctx context.Context, msize uint32) error {
	f := &Fcall{
		Args:  &proto.Tversion{Version: version, Msize: msize},
		Reply: &proto.Rversion{},
	}
	ch := make(chan *Fcall, 1)

	go c.Do(ctx, f, ch)
	f = <-ch
	log.Printf("-> %+v", f.Reply)

	return f.Err()
}
