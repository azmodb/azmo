package rpc

import (
	"context"
	"io"
	"net"
	"sync"

	protobuf "github.com/azmodb/protobuf"
	"github.com/azmodb/rpc/pb"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Client struct {
	writer sync.Mutex // exclusive writer lock
	enc    *protobuf.Encoder
	req    pb.Request
	c      io.Closer

	dec *protobuf.Decoder

	mu       sync.Mutex // protects following
	seq      uint64
	pending  map[uint64]*call
	closing  bool
	shutdown bool
}

func Dial(ctx context.Context, network, addr string) (*Client, error) {
	conn, err := (&net.Dialer{}).DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}
	return NewClient(conn, 0), nil
}

func NewClient(rwc io.ReadWriteCloser, max int) *Client {
	c := &Client{
		enc:     protobuf.NewEncoder(rwc, max),
		dec:     protobuf.NewDecoder(rwc, max),
		c:       rwc,
		seq:     1,
		pending: make(map[uint64]*call),
	}
	go c.receive()
	return c
}

type Event struct {
	Args  proto.Message
	Reply proto.Message
	Err   error
}

type call struct {
	ev *Event
	ch chan *Event
}

func (c *call) done() {
	select {
	case c.ch <- c.ev:
	default:
		// TODO: silently discard message; should not happen!
	}
}

func (c *Client) send(call *call) {
	c.mu.Lock()
	if c.shutdown || c.closing {
		call.ev.Err = errors.New("client is shut down")
		c.mu.Unlock()
		call.done()
		return
	}
	seq := c.seq
	c.seq++
	c.pending[seq] = call
	c.mu.Unlock()

	c.req.Method = pb.FlowMethod
	c.req.Seq = seq
	if err := c.encode(&c.req, call.ev.Args); err != nil {
		c.mu.Lock()
		call = c.pending[seq]
		delete(c.pending, seq)
		c.mu.Unlock()
		if call != nil {
			call.ev.Err = err
			call.done()
		}
	}
}

func (c *Client) encode(header *pb.Request, m proto.Message) error {
	var err error
	if err = c.enc.Encode(header); err != nil {
		return err
	}
	if err = c.enc.Encode(m); err != nil {
		return err
	}
	return c.enc.Flush()
}

func (c *Client) receive() {
	var resp pb.Response
	var err error
	for err == nil {
		resp = pb.Response{}
		if err = c.dec.Decode(&resp); err != nil {
			break
		}

		c.mu.Lock()
		call := c.pending[resp.Seq]
		delete(c.pending, resp.Seq)
		c.mu.Unlock()

		switch {
		case call == nil:
			if err = c.dec.Decode(nil); err != nil {
				if err != io.EOF && err != io.ErrUnexpectedEOF {
					err = errors.Wrap(err, "decoding error body")
				}
			}
		case resp.Error != nil:
			call.ev.Err = resp.Error
			if err = c.dec.Decode(nil); err != nil {
				if err != io.EOF && err != io.ErrUnexpectedEOF {
					err = errors.Wrap(err, "decoding body")
				}
			}
			call.done()
		default:
			if err = c.dec.Decode(call.ev.Reply); err != nil {
				if err != io.EOF && err != io.ErrUnexpectedEOF {
					err = errors.Wrap(err, "decoding body")
				}
			}
			call.done()
		}
	}

	c.writer.Lock()
	c.mu.Lock()
	c.shutdown = true
	if err == io.EOF {
		if c.closing {
			err = errors.New("client is shut down")
		} else {
			err = io.ErrUnexpectedEOF
		}
	}
	for _, call := range c.pending {
		call.ev.Err = err
		call.done()
	}
	c.mu.Unlock()
	c.writer.Unlock()
}

func (c *Client) Call(ctx context.Context, ev *Event) <-chan *Event {
	call := &call{ev: ev, ch: make(chan *Event, 1)}
	c.writer.Lock()
	c.send(call)
	c.writer.Unlock()
	return call.ch
}

func (c *Client) Send(ctx context.Context, m proto.Message) error {
	c.writer.Lock()
	c.writer.Unlock()
	return nil
}

func (c *Client) Recv(ctx context.Context, m proto.Message) error {
	return nil
}

func (c *Client) Close() error {
	c.mu.Lock()
	if c.closing {
		c.mu.Unlock()
		return errors.New("client is shut down")
	}
	c.closing = true
	c.mu.Unlock()
	return c.c.Close()
}
