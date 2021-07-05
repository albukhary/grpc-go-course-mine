package main

import (
	"fmt"
	"log"
	"net"

	"github.com/albukhary/grpc-go-course-mine/greet/greetpb"
	
	"google.golang.org/grpc"
)

type server struct{}

func main(){
	fmt.Println("Salam")

	// binding port 
	lis, err := net.Listen("tcp","0.0.0.50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	// createa gRPC server
	s := grpc.NewServer()
	// register services 
	__greetpb.RegisterGreetServiceServer(s, &server{} )

	// bind the port to the gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}