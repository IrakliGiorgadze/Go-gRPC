package main

import (
	"context"
	"fmt"
	"log"

	"gRPC/greet/greetpb"

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
	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Shady",
			LastName:  "Claus",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling RPC: %v", err)
	}
	log.Printf("response calling RPC: %v", res.Result)
}
