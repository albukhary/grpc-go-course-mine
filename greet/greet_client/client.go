package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/albukhary/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("salam, I am client 0_0")

	// Open an INSECURE client connection(cc)
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client Could not connect %v\n", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	//	doUnary(c)
	doServerStream(c)

}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do unary PRC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Lazizbek",
			LastName:  "Kahramonov",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStream(c greetpb.GreetServiceClient) {
	fmt.Println("Doing server streaming RPC ...")

	// Form a GreetManyTimes Request to send it to server
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Lazizbek",
			LastName:  "Kahramonov",
		},
	}

	// send request and get the response
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling GreetManyTimes RPC ...")
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v\n", err)
		}
		log.Printf("Response from GreetManyTimes : %v\n", msg.Result)
	}
}
