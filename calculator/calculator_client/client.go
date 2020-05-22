package main

import (
	"context"
	"fmt"
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
	doUnary(c)

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
