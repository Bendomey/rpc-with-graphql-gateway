package user

import (
	"context"
	"fmt"
	"net"

	"github/Bendomey/peerstronix-store/user/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// grpc server
type grpcServer struct {
	service Service
}

//ListenToGrpc listen's to grpc
func ListenToGrpc(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterUserServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

// functions for necessary procedures should be available here
func (s *grpcServer) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	a, err := s.service.CreateUser(ctx, r.Name, r.Phone, r.Email, r.Password)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:        a.ID,
			Name:      a.Name,
			Phone:     a.Phone,
			Email:     a.Email,
			IsDeleted: a.IsDeleted,
		},
	}, err
}

func (s *grpcServer) UpdateUser(ctx context.Context, r *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	a, err := s.service.UpdateUser(ctx, r.Name, r.Phone, r.Email, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		User: &pb.User{
			Id:        a.ID,
			Name:      a.Name,
			Phone:     a.Phone,
			Email:     a.Email,
			IsDeleted: a.IsDeleted,
		},
	}, err
}

func (s *grpcServer) GetUsers(ctx context.Context, r *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	res, err := s.service.GetUsers(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	users := []*pb.User{}
	for _, p := range res {
		users = append(users, &pb.User{
			Id:        p.ID,
			Name:      p.Name,
			Phone:     p.Phone,
			Email:     p.Email,
			IsDeleted: p.IsDeleted,
		})
	}
	return &pb.GetUsersResponse{User: users}, nil
}

func (s *grpcServer) GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	a, err := s.service.GetUser(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:        a.ID,
			Email:     a.Email,
			IsDeleted: a.IsDeleted,
			Name:      a.Name,
			Password:  a.Password,
			Phone:     a.Password,
		},
	}, err
}

// DeleteUser to delete user form db
func (s *grpcServer) DeleteUser(ctx context.Context, r *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	a, err := s.service.DeleteUser(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{
		Done: a,
	}, err
}
