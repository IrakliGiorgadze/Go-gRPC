package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"gRPC/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hi " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hi " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetmanytimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request")
	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreatResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("error reading client stream: %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += firstName + "! "
	}
}

func (*server) GreatEveryone(stream greetpb.GreetService_GreatEveryoneServer) error {
	fmt.Printf("GreatEveryone function was invoked with a streaming request")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error reading client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hi " + firstName + "! "
		err = stream.Send(&greetpb.GreatEveryoneResponse{
			Result: result,
		})
		if err != nil {
			log.Fatalf("error sending client stream: %v", err)
			return err
		}
	}
}

func main() {
	fmt.Println("server is running...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("error connection: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
