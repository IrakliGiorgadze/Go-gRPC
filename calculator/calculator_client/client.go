package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"gRPC/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("client is running...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalcServiceClient(cc)
	
	//doUnary(c)
	doServeStreaming(c)
}

func doUnary(c calculatorpb.CalcServiceClient) {
	req := &calculatorpb.SumRequest{
		FirstNumber:  9,
		SecondNumber: 60,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling RPC: %v", err)
	}
	log.Printf("response calling RPC: %v", res.SumResult)
}

func doServeStreaming(c calculatorpb.CalcServiceClient) {
	fmt.Println("starting streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 12,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling streaming RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error calling streaming RPC: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}