package main

import (
	"context"
	"fmt"
	"log"

	"gRPC/blog/blogpb"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("client is running...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)
	blog := &blogpb.Blog{
		AuthorId: "Shady",
		Title:    "Qubit",
		Content:  "In quantum computing, a qubit or quantum bit is the basic unit of quantum information—the quantum version of the classical binary bit physically realized with a two-state device. A qubit is a two-state quantum-mechanical system, one of the simplest quantum systems displaying the peculiarity of quantum mechanics",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	fmt.Println(createBlogRes)
}
