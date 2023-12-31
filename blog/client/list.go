package main

import (
	"context"
	"io"
	"log"

	pb "github.com/viralparmarme/go-grpc-mongodb-blog/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("listBlog was invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("something happened: %v\n", err)
		}

		log.Println(res)
	}
}
