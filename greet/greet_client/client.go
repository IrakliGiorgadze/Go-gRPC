package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)

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

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting client streaming RPC...")
	requests := []*greetpb.LongGreatRequest{
		&greetpb.LongGreatRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Shady",
				LastName:  "Claus",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error calling client streaming: %v", err)
	}
	for _, req := range requests {
		fmt.Printf("sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving client streaming: %v", err)
	}
	fmt.Printf("LongGreat response: %v\n", res)
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting client BiDi streaming RPC...")
	stream, err := c.GreatEveryone(context.Background())
	if err != nil {
		log.Fatalf("error calling client streaming: %v", err)
		return
	}

	requests := []*greetpb.GreatEveryoneRequest{
		&greetpb.GreatEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Donald",
				LastName:  "Trump",
			},
		},
		&greetpb.GreatEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Barack",
				LastName:  "Obama",
			},
		},
		&greetpb.GreatEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Bill",
				LastName:  "Gates",
			},
		},
		&greetpb.GreatEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Steve",
				LastName:  "Jobs",
			},
		},
		&greetpb.GreatEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Elon",
				LastName:  "Mask",
			},
		},
	}

	waitC := make(chan struct{})

	go func() {
		for _, req := range requests {
			fmt.Printf("sending message: %v", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error receiving client streaming: %v", err)
				break
			}
			fmt.Printf("received: %v\n", res.GetResult())
		}
		close(waitC)
	}()

	<-waitC
}
