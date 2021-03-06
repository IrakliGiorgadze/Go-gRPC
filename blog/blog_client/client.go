package main

import (
	"context"
	"fmt"
	"io"
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

	// Create Blog

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
	blogID := createBlogRes.GetBlog().GetId()

	// Read Blog

	_, err = c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "5ecbb195229e4f84e926237c",
	})
	if err != nil {
		fmt.Printf("error while reading: %v", err)
	}

	readBlogReq := &blogpb.ReadBlogRequest{
		BlogId: blogID,
	}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("error while reading: %v", err)
	}
	fmt.Println(readBlogRes)

	// Update Blog

	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Shady Claus",
		Title:    "Qubit (edited)",
		Content:  "In quantum computing... (edited)",
	}

	_, err = c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: newBlog,
	})
	if err != nil {
		fmt.Printf("error while updating: %v", err)
	}

	// Delete Blog

	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogID})
	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", deleteErr)
	}
	fmt.Printf("Blog was deleted: %v \n", deleteRes)

	// List Blog

	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
