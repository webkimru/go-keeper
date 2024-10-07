package client

import (
	"fmt"
	"google.golang.org/grpc"
)

// Client is a new gRPC client.
type Client struct {
	Client *grpc.ClientConn
	//error  chan error
}

// NewClient returns a new gRPC client.
func NewClient(address string, option ...grpc.DialOption) (*Client, error) {
	conn, err := grpc.NewClient(address, option...)
	if err != nil {
		return nil, err
	}

	c := &Client{
		Client: conn,
		//error:  make(chan error, 1),
	}

	return c, nil
}

//// Notify returns an error.
//func (c *Client) Notify() <-chan error {
//	return c.error
//}

// Close closes a gRPC connection.
func (c *Client) Close() error {
	if err := c.Client.Close(); err != nil {
		return fmt.Errorf("grpcserver - Client - Close - c.Client.Close(): %w", err)
	}

	return nil
}
