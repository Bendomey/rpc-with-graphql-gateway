package user

import (
	"github/Bendomey/peerstronix-store/user/pb"

	"google.golang.org/grpc"
)

// Client for global types in this module
type Client struct {
	conn    *grpc.ClientConn
	service pb.UserServiceClient
}

// NewClient is to create a client for rpc server
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewUserServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close is used to close the client
func (c *Client) Close() {
	c.conn.Close()
}
