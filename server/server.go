package server

import (
	"errors"
	"log"
	"net"
	"os"

	"github.com/azmodb/azmo/pb"
	"github.com/azmodb/db"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	log *log.Logger
	db  *db.DB
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.Event, error) {
	s.printf("GET -> %s", req)
	rec, err := s.db.Get(req.Key, req.Rev, req.Versions)
	if err != nil {
		rec.Close()
		return nil, err
	}

	ev := &pb.Event{Record: rec.Record, Duration: 0}
	s.printf("GET <- %s", ev)
	return ev, nil
}

func (s *server) Batch(req *pb.BatchRequest, srv pb.DB_BatchServer) (err error) {
	batch := s.db.Next()
	ev := &pb.Event{}

	for _, arg := range req.GetArgs() {
		var rec *db.Record

		switch t := arg.Command.(type) {
		case *pb.Argument_Put:
			r := t.Put
			if !r.Tombstone {
				rec, err = batch.Insert(r.Key, r.Value, r.Prev)
			} else {
				rec, err = batch.Put(r.Key, r.Value, r.Prev)
			}
		case *pb.Argument_Delete:
			r := t.Delete
			rec, err = batch.Delete(r.Key, r.Prev)
		case *pb.Argument_Increment:
			r := t.Increment
			rec, err = batch.Increment(r.Key, r.Value, r.Prev)
		case *pb.Argument_Decrement:
			r := t.Decrement
			rec, err = batch.Decrement(r.Key, r.Value, r.Prev)
		default:
			err = errors.New("unknown batch command")
		}
		if err != nil { // rec is already released
			break
		}

		if rec != nil {
			ev.Record = rec.Record
		}
		err = srv.Send(ev)
		rec.Close()
		if err != nil {
			break
		}
	}

	if err != nil {
		batch.Rollback()
	} else {
		batch.Commit()
	}
	return err
}

func (s *server) Range(req *pb.RangeRequest, srv pb.DB_RangeServer) error {
	return nil
}

func (s *server) Watch(req *pb.WatchRequest, srv pb.DB_WatchServer) error {
	return nil
}

func (s *server) printf(format string, args ...interface{}) {
	if s.log != nil {
		s.log.Printf(format, args...)
	}
}

type Option func(*server) error

func WithLogger(logger *log.Logger) Option {
	return func(s *server) error {
		s.log = logger
		return nil
	}
}

func Listen(listener net.Listener, opts ...Option) error {
	server := &server{db: db.New()}
	for _, opt := range opts {
		if err := opt(server); err != nil {
			return err
		}
	}

	server.log = log.New(os.Stderr, "", 0)

	s := grpc.NewServer()
	pb.RegisterDBServer(s, server)
	return s.Serve(listener)
}
