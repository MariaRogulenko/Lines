package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/MariaRogulenko/lines"
	"github.com/MariaRogulenko/lines/api"
)

var (
	gRPCPort = flag.Int("grpc-port", 8090, "GRPC Server port")
	httpPort = flag.Int("http-port", 8080, "HTTP Server port")
)

func runServer(s *lines.Service) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *gRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	api.RegisterGameServer(srv, s)
	if err := srv.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func registerProxy(ctx context.Context) {
	mux := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	if err := api.RegisterGameHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", *gRPCPort), dialOpts); err != nil {
		log.Fatal(err)
	}
	http.Handle("/api/", mux)
}

func main() {
	flag.Parse()
	s := lines.NewService()
	go runServer(s)
	registerProxy(context.Background())
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil); err != nil {
		log.Fatal(err)
	}
}
