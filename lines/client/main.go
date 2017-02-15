package main

import (
	"fmt"
	"log"

	pb "lines/lines/lines"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUpdateClient(conn)

	// Contact the server and print out its response.

	r, err := c.UpdateTable(context.Background(), &pb.StatusRequest{X: 1, Y: 2})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf(" %d", r.Table[i*9+j])
		}
		fmt.Print("\n")
	}
	if r.Active {
		fmt.Printf("Active %d %d ", r.X, r.Y)
	}
	log.Println(r.Table[0])
	log.Printf("Greeting: %d %d\n", r.X, r.Y)
}
