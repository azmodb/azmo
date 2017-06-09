package rpc

import (
	"io"
	"sync"

	protobuf "github.com/azmodb/protobuf"
	"github.com/azmodb/rpc/pb"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

var (
	//reqPool  = &sync.Pool{New: func() interface{} { return &pb.Request{} }}
	respPool = &sync.Pool{New: func() interface{} { return &pb.Response{} }}
)

type Server struct {
	f Factory
}

func NewServer(factory Factory, opts ...Option) (*Server, error) {
	return nil, nil
}

type Factory interface {
	Run(req proto.Message) (resp proto.Message, err error)
	New() proto.Message
}

type codec struct {
	writer sync.Mutex // exclusive writer lock
	enc    *protobuf.Encoder

	dec *protobuf.Decoder
}

func (c *codec) encode(header *pb.Response, m proto.Message) error {
	var err error
	if err = c.enc.Encode(header); err != nil {
		return err
	}
	if err = c.enc.Encode(m); err != nil {
		return err
	}
	return c.enc.Flush()
}

func (c *codec) decode(m proto.Message) error {
	return c.dec.Decode(m)
}

func (s *Server) Serve(rwc io.ReadWriteCloser) (err error) {
	c := &codec{
		dec: protobuf.NewDecoder(rwc, 0),
		enc: protobuf.NewEncoder(rwc, 0),
	}
	req := pb.Request{}
	defer rwc.Close()

	for err == nil {
		req.Reset()
		if err = c.decode(&req); err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				return err
			}
			err = errors.Wrap(err, "cannot decode request")
			// TODO: send back error if possible
			return err
		}

		m := s.f.New()
		if err = c.decode(m); err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				return err
			}
		}

		go func(req proto.Message, f Factory) {
			resp := respPool.Get().(*pb.Response)
			defer func() {
				resp.Reset()
				respPool.Put(resp)
			}()

			m, err := f.Run(req)
			if err != nil {
				resp.Error = &pb.Error{
					Description: err.Error(),
					ErrCode:     -1, // TODO
				}
			}

			c.encode(resp, m)
		}(m, s.f)
	}
	return err
}
