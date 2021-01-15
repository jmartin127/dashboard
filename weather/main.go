package main

import (
	"context"
	"log"
	"net"

	pb "github.com/jmartin127/dashboard/weather/weather"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

type server struct {
	pb.UnimplementedWeatherServer
}

// TODO move gRPC functions to separate package
func (s *server) GetCurrentWeather(ctx context.Context, in *pb.GetCurrentWeatherRequest) (*pb.GetCurrentWeatherReply, error) {
	// TODO integrate with Google API
	return &pb.GetCurrentWeatherReply{TempFahrenheit: 39, PrecipitationPct: 1, HumidityPct: 38, WindMPH: 5}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	reflection.Register(s) // needed for reflection API on the gRPC server

	pb.RegisterWeatherServer(s, &server{})

	log.Printf("Listening on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
