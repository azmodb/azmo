package azmo

import (
	"errors"
	"net"

	"github.com/azmodb/db"
	pb "github.com/azmodb/exp/azmo/azmopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	db *db.DB
}

func (s *server) Decrement(ctx context.Context, req *pb.DecrementRequest) (*pb.Event, error) {
	ev := &pb.Event{}
	tx := s.db.Txn()
	if err := s.decrement(tx, req, ev); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return ev, nil
}

func (s *server) Increment(ctx context.Context, req *pb.IncrementRequest) (*pb.Event, error) {
	ev := &pb.Event{}
	tx := s.db.Txn()
	if err := s.increment(tx, req, ev); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return ev, nil
}

func (s *server) decrement(tx *db.Txn, req *pb.DecrementRequest, ev *pb.Event) error {
	up := func(data interface{}) interface{} {
		v := data.(int64)
		return v - req.Value
	}

	created, err := tx.Update(req.Key, up, req.Tombstone)
	if err != nil {
		return err
	}

	ev.Init(pb.Decrement, req.Key, nil, created, created)
	return nil
}

func (s *server) increment(tx *db.Txn, req *pb.IncrementRequest, ev *pb.Event) error {
	up := func(data interface{}) interface{} {
		v := data.(int64)
		return v + req.Value
	}

	created, err := tx.Update(req.Key, up, req.Tombstone)
	if err != nil {
		return err
	}

	ev.Init(pb.Decrement, req.Key, nil, created, created)
	return nil
}

func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.Event, error) {
	ev := &pb.Event{}
	tx := s.db.Txn()
	if err := s.put(tx, req, ev); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return ev, nil
}

func (s *server) put(tx *db.Txn, req *pb.PutRequest, ev *pb.Event) error {
	created, err := tx.Put(req.Key, req.Value, req.Tombstone)
	if err != nil {
		return err
	}

	ev.Init(pb.Put, req.Key, nil, created, created)
	return nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.Event, error) {
	ev := &pb.Event{}
	tx := s.db.Txn()
	if err := s.delete(tx, req, ev); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return ev, nil
}

func (s *server) delete(tx *db.Txn, req *pb.DeleteRequest, ev *pb.Event) error {
	current := tx.Delete(req.Key)

	ev.Init(pb.Delete, req.Key, nil, current, current)
	return nil
}

func (s *server) Batch(req *pb.BatchRequest, srv pb.DB_BatchServer) (err error) {
	ev := &pb.Event{}
	tx := s.db.Txn()
	for _, arg := range req.Arguments {
		switch t := arg.Command.(type) {
		case *pb.Argument_Decrement:
			err = s.decrement(tx, t.Decrement, ev)
		case *pb.Argument_Increment:
			err = s.increment(tx, t.Increment, ev)
		case *pb.Argument_Put:
			err = s.put(tx, t.Put, ev)
		case *pb.Argument_Delete:
			err = s.delete(tx, t.Delete, ev)
		default:
			err = errors.New("unsupported batch argument")
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.Event, error) {
	value, created, current, err := s.db.Get(req.Key, req.Revision, req.MustEqual)
	if err != nil {
		return nil, err
	}
	return pb.NewEvent(pb.Get, req.Key, value, created, current), nil
}

func (s *server) Range(req *pb.RangeRequest, srv pb.DB_RangeServer) error {
	n, _, err := s.db.Range(req.From, req.To, req.Revision, req.Limit)
	if err != nil {
		return err
	}
	defer n.Cancel()

	ev := &pb.Event{}
	for e := range n.Recv() {
		evErr := e.Err()
		if evErr != nil && evErr == db.NotifierCanceled {
			return nil
		}
		if evErr != nil {
			return evErr
		}

		ev.Init(pb.Range, e.Key, e.Data, e.Created, e.Current)
		if err = srv.Send(ev); err != nil {
			return err
		}
	}
	return err
}

func (s *server) Watch(req *pb.WatchRequest, srv pb.DB_WatchServer) error {
	n, _, err := s.db.Watch(req.Key)
	if err != nil {
		return err
	}
	defer n.Cancel()

	ev := &pb.Event{}
	for e := range n.Recv() {
		evErr := e.Err()
		if evErr != nil && evErr == db.NotifierCanceled {
			return nil
		}
		if evErr != nil {
			return evErr
		}

		ev.Init(pb.Watch, e.Key, e.Data, e.Created, e.Current)
		if err = srv.Send(ev); err != nil {
			return err
		}
	}
	return err
}

type ServerOption func(*server) error

func Listen(db *db.DB, listener net.Listener, cert, key string, opts ...ServerOption) (err error) {
	server := &server{db: db}
	for _, opt := range opts {
		if err = opt(server); err != nil {
			return err
		}
	}

	var options []grpc.ServerOption
	if cert != "" && key != "" {
		creds, err := credentials.NewServerTLSFromFile(cert, key)
		if err != nil {
			return err
		}
		options = []grpc.ServerOption{grpc.Creds(creds)}
	}

	s := grpc.NewServer(options...)
	pb.RegisterDBServer(s, server)
	return s.Serve(listener)
}
