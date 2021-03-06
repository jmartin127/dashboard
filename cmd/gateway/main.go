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
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	dashboardgw "github.com/jmartin127/dashboard/proto/gen/go/jmartin127/dashboard/v1"
	trafficgw "github.com/jmartin127/dashboard/proto/gen/go/jmartin127/traffic/v1"
	usergw "github.com/jmartin127/dashboard/proto/gen/go/jmartin127/user/v1"
	weathergw "github.com/jmartin127/dashboard/proto/gen/go/jmartin127/weather/v1"
)

var (
	grpcServerEndpointTraffic   = flag.String("grpc-server-endpoint-traffic", "localhost:50051", "gRPC Traffic server endpoint")
	grpcServerEndpointWeather   = flag.String("grpc-server-endpoint-weather", "localhost:50052", "gRPC Weather server endpoint")
	grpcServerEndpointDashboard = flag.String("grpc-server-endpoint-dashboard", "localhost:50053", "gRPC Dashboard server endpoint")
	grpcServerEndpointUser      = flag.String("grpc-server-endpoint-user", "localhost:50054", "gRPC User server endpoint")
)

const (
	port = ":8081"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// set up regular net/http mux
	mux := http.NewServeMux()

	// Register gRPC server endpoint
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register the handlers for each service
	if err := trafficgw.RegisterTrafficHandlerFromEndpoint(ctx, gwmux, *grpcServerEndpointTraffic, opts); err != nil {
		return err
	}
	if err := weathergw.RegisterWeatherHandlerFromEndpoint(ctx, gwmux, *grpcServerEndpointWeather, opts); err != nil {
		return err
	}
	if err := dashboardgw.RegisterDashboardHandlerFromEndpoint(ctx, gwmux, *grpcServerEndpointDashboard, opts); err != nil {
		return err
	}
	if err := usergw.RegisterUsersHandlerFromEndpoint(ctx, gwmux, *grpcServerEndpointUser, opts); err != nil {
		return err
	}

	// set up swagger
	mux.Handle("/", gwmux)
	serveSwagger(mux)

	log.Printf("Gateway listening on port %s\n", port)
	return http.ListenAndServe(port, mux)
}

func serveSwagger(mux *http.ServeMux) {
	swaggerDir := os.Getenv("SWAGGER_UI_DIR")
	fs := http.FileServer(http.Dir(swaggerDir))
	mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
