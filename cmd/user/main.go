package main

import (
	"context"
	"log"
	"net"

	empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/jmartin127/dashboard/proto/gen/go/jmartin127/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port = ":50054"
)

var users = map[string]*pb.User{}

type server struct {
	pb.UnimplementedUsersServer
}

func (s *server) AddUser(ctx context.Context, in *pb.AddUserRequest) (*empty.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// check if the user already exists
	if _, ok := users[in.User.Username]; ok {
		return nil, status.Error(codes.AlreadyExists, "This user already exists")
	}

	// add the user
	users[in.User.Username] = in.User

	return &empty.Empty{}, nil
}

func (s *server) ModifyUser(ctx context.Context, in *pb.ModifyUserRequest) (*empty.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// make sure the user exists
	if _, ok := users[in.User.Username]; !ok {
		return nil, status.Error(codes.NotFound, "This user was not found")
	}

	// update the user
	users[in.User.Username] = in.User

	return &empty.Empty{}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserReply, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// make sure the user exists
	if _, ok := users[in.Username]; !ok {
		return nil, status.Error(codes.NotFound, "This user was not found")
	}

	// return the user
	return &pb.GetUserReply{User: users[in.Username]}, nil
}

func (s *server) RemoveUser(ctx context.Context, in *pb.RemoveUserRequest) (*empty.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// make sure the user exists
	if _, ok := users[in.Username]; !ok {
		return nil, status.Error(codes.NotFound, "This user was not found")
	}

	// remove the user
	delete(users, in.Username)

	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	reflection.Register(s) // needed for reflection API on the gRPC server

	pb.RegisterUsersServer(s, &server{})

	log.Printf("User Service listening on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
