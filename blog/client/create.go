package main

import (
	"context"
	"log"

	pb "github.com/viralparmarme/go-grpc-mongodb-blog/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Viral",
		Title:    "First Blog",
		Content:  "Content comes here",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("unexpected error %v\n", err)
	}

	log.Printf("blog has been created: %s\n", res.Id)
	return res.Id
}
