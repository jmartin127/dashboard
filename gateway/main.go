package main

import (
	// +build tools
	// _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	// _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	// _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	// _ "google.golang.org/protobuf/cmd/protoc-gen-go"

	"context"
	"flag"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	trafficgw "github.com/jmartin127/dashboard/proto/gen/go/traffic"
	weathergw "github.com/jmartin127/dashboard/proto/gen/go/weather"
)

var (
	grpcServerEndpointTraffic = flag.String("grpc-server-endpoint-traffic", "localhost:50051", "gRPC Traffic server endpoint")
	grpcServerEndpointWeather = flag.String("grpc-server-endpoint-weather", "localhost:50052", "gRPC Weather server endpoint")
)

const (
	port = ":8081"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register the handlers for each service
	if err := trafficgw.RegisterTrafficHandlerFromEndpoint(ctx, mux, *grpcServerEndpointTraffic, opts); err != nil {
		return err
	}
	if err := weathergw.RegisterWeatherHandlerFromEndpoint(ctx, mux, *grpcServerEndpointWeather, opts); err != nil {
		return err
	}

	log.Printf("Listening on port %s\n", port)
	return http.ListenAndServe(port, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
