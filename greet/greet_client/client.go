package main

import (
	"fmt"
	"gRPC/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("client is running...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("created client: %f", c)
}
