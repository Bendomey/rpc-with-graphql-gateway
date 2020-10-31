package user

import (
	"context"
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

// CreateUser for creating a userr
func (c *Client) CreateUser(ctx context.Context, name string, phone string, email string, password string) (*User, error) {
	a, err := c.service.CreateUser(ctx, &pb.CreateUserRequest{Name: name, Email: email, Password: password, Phone: phone})
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    a.User.Id,
		Name:  a.User.Name,
		Phone: a.User.Phone,
		Email: a.User.Email,
	}, err
}

// UpdateUser for updating a userr
func (c *Client) UpdateUser(ctx context.Context, name string, phone string, email string, password string, id string) (*User, error) {
	a, err := c.service.UpdateUser(ctx, &pb.UpdateUserRequest{Name: name, Email: email, Phone: phone, Id: id})
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    a.User.Id,
		Name:  a.User.Name,
		Phone: a.User.Phone,
		Email: a.User.Email,
	}, err
}

// GetUsers for getting users on a criteria
func (c *Client) GetUsers(ctx context.Context, skip uint64, take uint64) ([]User, error) {
	a, err := c.service.GetUsers(ctx, &pb.GetUsersRequest{Skip: skip, Take: take})
	if err != nil {
		return nil, err
	}

	users := []User{}
	for _, u := range a.User {
		users = append(users, User{
			ID:    u.Id,
			Name:  u.Name,
			Email: u.Email,
			Phone: u.Phone,
		})
	}

	return users, err
}

// GetUser for fethcing a single userr
func (c *Client) GetUser(ctx context.Context, id string) (*User, error) {
	a, err := c.service.GetUser(ctx, &pb.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    a.User.Id,
		Name:  a.User.Name,
		Phone: a.User.Phone,
		Email: a.User.Email,
	}, err
}

// DeleteUser for deleting a userr
func (c *Client) DeleteUser(ctx context.Context, id string) (bool, error) {
	a, err := c.service.DeleteUser(ctx, &pb.DeleteUserRequest{Id: id})
	if err != nil {
		return false, err
	}

	return a.Done, err
}
