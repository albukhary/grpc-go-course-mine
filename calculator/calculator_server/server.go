package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/albukhary/grpc-go-course-mine/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSumServiceServer
}

func (*server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {

	// Display the sum request from the client
	fmt.Printf("Server accepted Sum request: %v", req)

	// read the request to find values of number 1 and number 2
	result := req.GetNum1() + req.GetNum2()

	// Create a SumResponse
	res := &pb.SumResponse{
		Sum: result,
	}

	// return the response
	return res, nil
}

func main() {
	// Create a network listener, bind it to port 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()

	// Register a service to the gRPC server
	pb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
