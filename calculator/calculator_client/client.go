package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	//doServeStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)
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

func doClientStreaming(c calculatorpb.CalcServiceClient) {
	fmt.Println("starting doClientStreaming streaming RPC...")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error calling ComputeAverage streaming RPC: %v", err)
	}
	numbers := []int32{3, 5, 9, 54, 23}
	for _, number := range numbers {
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving ComputeAverage response RPC: %v", err)
	}
	fmt.Printf("The Average is: %v", res.GetAverage())
}

func doBiDiStreaming(c calculatorpb.CalcServiceClient) {
	fmt.Println("starting doBiDiStreaming streaming RPC...")

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("error opening and calling FindMaximum: %v", err)
		return
	}

	waitC := make(chan struct{})

	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: number,
			})
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
				log.Fatalf("error reading server stream: %v", err)
				break
			}
			maximum := res.GetMaximum()
			fmt.Printf("recieved a new maximum of: %v", maximum)
		}
		close(waitC)
	}()

	<-waitC
}
