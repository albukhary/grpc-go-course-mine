package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/albukhary/grpc-go-course-mine/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Salam I am calculator client")

	fmt.Println("salam, I am client 0_0")

	// Open an INSECURE client connection(cc)
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client Could not connect %v\n", err)
	}
	defer cc.Close()

	c := pb.NewSumServiceClient(cc)
	doUnary(c)
}

func doUnary(c pb.SumServiceClient) {
	fmt.Println("Starting to do unary PRC...")

	req := &pb.SumRequest{
		Num1: 4,
		Num2: 5,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum = %v\n", res.Sum)
}
