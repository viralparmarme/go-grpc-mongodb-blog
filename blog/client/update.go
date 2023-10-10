package main

import (
	"context"
	"log"

	pb "github.com/viralparmarme/go-grpc-mongodb-blog/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Viral",
		Title:    "New Title",
		Content:  "Updated content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("error happened during update: %v\n", err)
	}

	log.Println("blog was updated")
}
