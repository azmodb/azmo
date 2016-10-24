package azmo

import (
	"errors"
	"fmt"
	"io"
	"time"

	pb "github.com/azmodb/azmo/azmopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ClientOption func(*Client) error

type Client struct {
	conn *grpc.ClientConn
	db   pb.DBClient
}

func Dial(address string, timeout time.Duration, opts ...ClientOption) (*Client, error) {
	options := []grpc.DialOption{grpc.WithInsecure()}
	if timeout > 0 {
		options = append(options, grpc.WithTimeout(timeout))
	}
	options = append(options, grpc.WithBlock())

	conn, err := grpc.Dial(address, options...)
	if err != nil {
		return nil, err
	}
	return NewClient(conn), nil
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{db: pb.NewDBClient(conn), conn: conn}
}

func (c *Client) Close() error {
	if c == nil || c.conn == nil || c.db == nil {
		return errors.New("database connection is shut down")
	}

	err := c.conn.Close()
	c.conn = nil
	c.db = nil
	return err
}

type Event struct {
	Duration time.Duration
	*pb.Event
}

func newEvent(start time.Time, resp *pb.Event) *Event {
	return &Event{
		Duration: time.Now().Sub(start),
		Event:    resp,
	}
}

func (e *Event) String() string {
	return fmt.Sprintf("%sduration:%s", e.Event, e.Duration)
}

func (c *Client) Batch(ctx context.Context, req *pb.BatchRequest) ([]*Event, error) {
	stream, err := c.db.Batch(ctx, req)
	if err != nil {
		return nil, err
	}

	evs := make([]*Event, 0, len(req.Arguments))
	for {
		start := time.Now()
		resp, err := stream.Recv()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		evs = append(evs, newEvent(start, resp))
	}
	return evs, nil
}

func (c *Client) Delete(ctx context.Context, req *pb.DeleteRequest) (*Event, error) {
	start := time.Now()
	resp, err := c.db.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return newEvent(start, resp), nil
}

func (c *Client) Dec(ctx context.Context, req *pb.DecrementRequest) (*Event, error) {
	start := time.Now()
	resp, err := c.db.Decrement(ctx, req)
	if err != nil {
		return nil, err
	}
	return newEvent(start, resp), nil
}

func (c *Client) Inc(ctx context.Context, req *pb.IncrementRequest) (*Event, error) {
	start := time.Now()
	resp, err := c.db.Increment(ctx, req)
	if err != nil {
		return nil, err
	}
	return newEvent(start, resp), nil
}

func (c *Client) Put(ctx context.Context, req *pb.PutRequest) (*Event, error) {
	start := time.Now()
	resp, err := c.db.Put(ctx, req)
	if err != nil {
		return nil, err
	}
	return newEvent(start, resp), nil
}

func (c *Client) Range(ctx context.Context, req *pb.RangeRequest) ([]*Event, error) {
	stream, err := c.db.Range(ctx, req)
	if err != nil {
		return nil, err
	}

	limit := int32(64)
	if req.Limit != 0 {
		limit = req.Limit
	}

	evs := make([]*Event, 0, limit)
	for {
		start := time.Now()
		resp, err := stream.Recv()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		evs = append(evs, newEvent(start, resp))
	}
	return evs, nil
}

func (c *Client) Get(ctx context.Context, req *pb.GetRequest) (*Event, error) {
	start := time.Now()
	resp, err := c.db.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return newEvent(start, resp), nil
}

type Notifier interface {
	Notify(ev *Event) error
}

func (c *Client) Watch(ctx context.Context, req *pb.WatchRequest, n Notifier) error {
	stream, err := c.db.Watch(ctx, req)
	if err != nil {
		return err
	}

	for {
		start := time.Now()
		resp, err := stream.Recv()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err = n.Notify(newEvent(start, resp)); err != nil {
			return err
		}
	}
	return nil
}
