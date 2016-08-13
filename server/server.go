package server

import (
	"fmt"
	"log"
	"net"

	"github.com/azmodb/azmo/pb"
	"github.com/azmodb/db"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	logger *log.Logger
	db     *db.DB
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	s.trace("DELETE %s", req)

	resp := &pb.DeleteResponse{}
	txn := s.db.Txn()
	resp.Revs, resp.Rev = txn.Delete(req.Key)
	txn.Commit()

	s.end("DELETE %s", resp)
	return resp, nil
}

func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	s.trace("PUT %s", req)

	resp := &pb.PutResponse{}
	txn := s.db.Txn()
	resp.Revs, resp.Rev = txn.Put(req.Key, req.Value, req.Tombstone)
	txn.Commit()

	s.end("PUT %s", resp)
	return resp, nil
}

func (s *server) Txn(ctx context.Context, req *pb.TxnRequest) (*pb.TxnResponse, error) {
	s.trace("TXN %s", req)

	if len(req.Requests) == 0 {
		return nil, fmt.Errorf("txn requires at least 1 generic request")
	}

	resp := &pb.TxnResponse{
		Responses: make([]*pb.GenericResponse, 0, len(req.Requests)),
	}
	defer s.end("TXN %s", resp)

	txn := s.db.Txn()
	for _, r := range req.Requests {
		if r.Num <= 0 {
			txn.Rollback()
			return nil, fmt.Errorf("invalid txnid %d", r.Num)
		}

		t := &pb.GenericResponse{Num: r.Num}
		switch r.Type {
		case pb.GenericRequest_PutRequest:
			t.Revs, t.Rev = txn.Put(r.Key, r.Value, r.Tombstone)
		case pb.GenericRequest_DeleteRequest:
			t.Revs, t.Rev = txn.Delete(r.Key)
		default:
			txn.Rollback()
			return nil, fmt.Errorf("invalid request type %d", r.Type)
		}
		resp.Responses = append(resp.Responses, t)
	}
	if len(resp.Responses) != len(req.Requests) {
		txn.Rollback()
		return nil, fmt.Errorf("malformed responses length")
	}
	txn.Commit()

	resp.Responses = resp.Responses[:len(req.Requests)]
	return resp, nil
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	s.trace("GET %s", req)

	resp := &pb.GetResponse{}
	resp.Value, resp.Revs, resp.Rev = s.db.Get(req.Key, req.Rev)

	s.end("GET %s", resp)
	return resp, nil
}

func (s *server) Range(req *pb.RangeRequest, srv pb.DB_RangeServer) (err error) {
	s.trace("RANGE %s", req)

	resp := &pb.RangeResponse{}
	s.db.Range(req.From, req.To, req.Rev, func(k, v []byte, revs []int64, rev int64) bool {
		resp.Key = k
		resp.Value = v
		resp.Revs = revs
		resp.Rev = rev
		if err = srv.Send(resp); err != nil {
			return true
		}
		s.end("RANGE %s", resp)
		return false
	})
	return err
}

func (s *server) trace(format string, args ...interface{}) {
	if s.logger == nil {
		return
	}
	s.logger.Printf("-> "+format, args...)
}

func (s *server) end(format string, args ...interface{}) {
	if s.logger == nil {
		return
	}
	s.logger.Printf("<- "+format, args...)
}

type Option func(*server) error

func WithLogger(logger *log.Logger) Option {
	return func(s *server) error {
		s.logger = logger
		return nil
	}
}

func Listen(listener net.Listener, options ...Option) error {
	server := &server{db: db.New()}
	for _, opt := range options {
		if err := opt(server); err != nil {
			return err
		}
	}

	s := grpc.NewServer()
	pb.RegisterDBServer(s, server)
	return s.Serve(listener)
}
