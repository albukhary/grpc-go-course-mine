package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

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

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Println("GreetManyTimes function incoked")

	firstName := req.Greeting.FirstName

	for i := 0; i < 10; i++ {
		result := fmt.Sprintln("Salam", firstName, "number:", strconv.Itoa(i+1))

		// Form a GreetManyTimes response
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}

		// Send the result to the client
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
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
