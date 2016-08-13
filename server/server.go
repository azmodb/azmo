package server

import (
	"fmt"
	"net"

	"github.com/azmodb/azmo/pb"
	"github.com/azmodb/db"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	db *db.DB
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	resp := &pb.DeleteResponse{}
	txn := s.db.Txn()
	resp.Rev = txn.Delete(req.Key)
	txn.Commit()
	return resp, nil
}

func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	resp := &pb.PutResponse{}
	txn := s.db.Txn()
	resp.Rev = txn.Put(req.Key, req.Value, req.Tombstone)
	txn.Commit()
	return resp, nil
}

func (s *server) Txn(ctx context.Context, req *pb.TxnRequest) (*pb.TxnResponse, error) {
	if len(req.Requests) == 0 {
		return nil, fmt.Errorf("txn requires at least 1 request")
	}

	resp := &pb.TxnResponse{
		Responses: make([]*pb.GenericResponse, 0, len(req.Requests)),
	}
	txn := s.db.Txn()
	for _, r := range req.Requests {
		if r.Num <= 0 {
			txn.Rollback()
			return nil, fmt.Errorf("invalid txnid %d", r.Num)
		}

		t := &pb.GenericResponse{Num: r.Num}
		switch r.Type {
		case pb.GenericRequest_DeleteRequest:
			t.Rev = txn.Delete(r.Key)
		case pb.GenericRequest_PutRequest:
			t.Rev = txn.Put(r.Key, r.Value, r.Tombstone)
		default:
			txn.Rollback()
			return nil, fmt.Errorf("invalid request type %d", r.Type)
		}
		resp.Responses = append(resp.Responses, t)
	}
	if len(resp.Responses) != len(req.Requests) {
		txn.Rollback()
		return nil, fmt.Errorf("malformed response")
	}
	txn.Commit()

	resp.Responses = resp.Responses[:len(req.Requests)]
	return resp, nil
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	resp := &pb.GetResponse{}
	resp.Value, resp.Revs, resp.Rev = s.db.Get(req.Key, req.Rev)
	return resp, nil
}

func (s *server) Range(req *pb.RangeRequest, srv pb.DB_RangeServer) (err error) {
	s.db.Range(req.From, req.To, req.Rev, func(k, v []byte, revs []int64, rev int64) bool {
		resp := &pb.RangeResponse{Key: k, Value: v, Revs: revs, Rev: rev}
		if err = srv.Send(resp); err != nil {
			return true
		}
		return false
	})
	return err
}

func Listen(listener net.Listener) error {
	s := grpc.NewServer()
	pb.RegisterDBServer(s, &server{db: db.New()})
	return s.Serve(listener)
}
