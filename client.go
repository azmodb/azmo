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
	return newClient(conn), nil
}

func newClient(conn *grpc.ClientConn) *Client {
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

func (c *Client) Batch(ctx context.Context, w io.Writer, r *pb.BatchRequest) error {
	stream, err := c.db.Batch(ctx, r)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if _, err = fmt.Fprintln(w, resp); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) Delete(ctx context.Context, w io.Writer, r *pb.DeleteRequest) error {
	resp, err := c.db.Delete(ctx, r)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(w, resp)
	return err
}

func (c *Client) Put(ctx context.Context, w io.Writer, r *pb.PutRequest) error {
	resp, err := c.db.Put(ctx, r)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(w, resp)
	return err
}

func (c *Client) Range(ctx context.Context, w io.Writer, r *pb.RangeRequest) error {
	stream, err := c.db.Range(ctx, r)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if _, err = fmt.Fprintln(w, resp); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) Get(ctx context.Context, w io.Writer, r *pb.GetRequest) error {
	resp, err := c.db.Get(ctx, r)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(w, resp)
	return err
}

func (c *Client) Watch(ctx context.Context, w io.Writer, r *pb.WatchRequest) error {
	stream, err := c.db.Watch(ctx, r)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if _, err = fmt.Fprintln(w, resp); err != nil {
			return err
		}
	}
	return nil
}
