package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"gRPC/blog/blogpb"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("blog server started...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("error connection: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("stopping the server")
	s.Stop()
	fmt.Println("closing the listener")
	lis.Close()
	fmt.Println("end of the program")
}
