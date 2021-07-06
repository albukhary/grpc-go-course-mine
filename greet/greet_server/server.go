package main

import (
	"context"
	"fmt"
	"log"
	"net"

	greetpb "github.com/albukhary/grpc-go-course-mine/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Printf("Greet function was invoked with request %v\n", req)

	// Extract information from our Request Object
	firstName := req.GetGreeting().GetFirstName()

	// Now form a greet response
	result := "Assalamu alaykum " + firstName

	// create a protocol buffer struct for output
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Salam")

	// Create a network listener, bind it to port 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()

	// Register a service to the gRPC server
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
