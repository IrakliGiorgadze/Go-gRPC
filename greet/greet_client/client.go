package main

import (
	"context"
	"fmt"
	"io"
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

	//doUnary(c)
	doServerStreaming(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting unary RPC...")
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

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting server streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Shady",
			LastName:  "Claus",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling server streaming: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error reading stream: %v", err)
		}
		log.Printf("response from GreetManyTimes: %v", msg.GetResult())
	}
}
