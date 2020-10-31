package account

import (
	"fmt"
	"net"
	pb "github.com/Bendomey/peerstronix-store/user/pb"
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
		return error
	}

	serv := grpc.NewServer()
	pb.RegisterAccountServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}
