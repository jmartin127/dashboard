package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/jmartin127/dashboard/traffic/traffic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedTrafficServer
}

func (s *server) GetTravelTime(ctx context.Context, in *pb.GetTravelTimeRequest) (*pb.GetTravelTimeReply, error) {
	// TODO integrate with Google API
	return &pb.GetTravelTimeReply{TravelTime: durationpb.New(11 * time.Minute)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	reflection.Register(s) // needed for reflection API on the gRPC server

	pb.RegisterTrafficServer(s, &server{})

	log.Printf("Listening on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
