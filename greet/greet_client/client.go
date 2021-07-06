package main

import (
	"fmt"
	"log"

	"github.com/albukhary/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main(){
	fmt.Println("salam, I am client 0_0")

	// Open an INSECURE client connection(cc)
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client Could not connect %v\n", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)
}