package main

import (
	"log"
	"net"

	pb "lines/lines/lines"

	"lines"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	currentStatus *pb.StatusReply
}

// SayHello implements helloworld.GreeterServer
func (s *server) UpdateTable(ctx context.Context, in *pb.StatusRequest) (*pb.StatusReply, error) {

	return &pb.StatusReply{Table: []int32{1, 2, 3, 1, 2, 3, 1, 2, 3,
		1, 2, 3, 1, 2, 3, 1, 2, 3,
		1, 2, 3, 1, 2, 3, 1, 2, 3,
		1, 2, 3, 1, 2, 3, 1, 2, 3,
		1, 2, 3, 1, 2, 3, 1, 2, 3,
		1, 2, 3, 1, 2, 3, 1, 2, 3,
		1, 2, 3, 1, 2, 3, 1, 2, 3,
		1, 2, 3, 1, 2, 3, 1, 2, 3,
		1, 2, 3, 1, 2, 3, 1, 2, 3}, Active: true, X: 1, Y: 2}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	q.PrintMatr()
	pb.RegisterUpdateServer(s, &server{})
	// Register reflection service on gRPC server.
	// reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
