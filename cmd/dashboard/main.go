package main

import (
	"context"
	"log"
	"net"

	pb "github.com/jmartin127/dashboard/proto/gen/go/jmartin127/dashboard/v1"
	trafficpb "github.com/jmartin127/dashboard/proto/gen/go/jmartin127/traffic/v1"
	weatherpb "github.com/jmartin127/dashboard/proto/gen/go/jmartin127/weather/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port = ":50053"
)

type server struct {
	pb.UnimplementedDashboardServer
	trafficClient trafficpb.TrafficClient
	weatherClient weatherpb.WeatherClient
}

func (s *server) GetDashboard(ctx context.Context, in *pb.GetDashboardRequest) (*pb.GetDashboardReply, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	travelTime, err := s.trafficClient.GetTravelTime(ctx, &trafficpb.GetTravelTimeRequest{OriginAddress: in.OriginAddress, DestinationAddress: in.DestinationAddress})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	weather, err := s.weatherClient.GetCurrentWeather(ctx, &weatherpb.GetCurrentWeatherRequest{Address: in.DestinationAddress})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetDashboardReply{
		Traffic: travelTime,
		Weather: weather,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s) // needed for reflection API on the gRPC server

	trafficClient, trafficConn, err := newTrafficClient()
	if err != nil {
		log.Fatalf("failed to create traffic client: %v", err)
	}
	defer trafficConn.Close()

	weatherClient, weatherConn, err := newWeatherClient()
	if err != nil {
		log.Fatalf("failed to create weather client: %v", err)
	}
	defer weatherConn.Close()

	pb.RegisterDashboardServer(s, &server{trafficClient: trafficClient, weatherClient: weatherClient})

	log.Printf("Dashboard Service listening on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func newTrafficClient() (trafficpb.TrafficClient, *grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		return nil, nil, err
	}

	client := trafficpb.NewTrafficClient(conn)
	return client, conn, nil
}

func newWeatherClient() (weatherpb.WeatherClient, *grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:50052", opts...)
	if err != nil {
		return nil, nil, err
	}

	client := weatherpb.NewWeatherClient(conn)
	return client, conn, nil
}
