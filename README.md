

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)


# Blog platform created using Go, gRPC, and MongoDB

Just another side project, this time it's a blog platform created using gRPC and MongoDB, all written in Go.


## Documentation

This documentation provides an overview of the Blog project. The project implements Create, Read, Update, and Delete (CRUD) operations for a blog service.

## Prerequisites

Before getting started, ensure that you have the following prerequisites installed on your system:

- Go (Golang)
- gRPC for Go
- MongoDB

## Project Structure

The project has the following structure:

```
├── Makefile
├── blog
│   ├── client
│   │   ├── create.go
│   │   ├── delete.go
│   │   ├── list.go
│   │   ├── main.go
│   │   ├── read.go
│   │   └── update.go
│   ├── docker-compose.yml
│   ├── proto
│   │   ├── blog.pb.go
│   │   ├── blog.proto
│   │   └── blog_grpc.pb.go
│   └── server
│       ├── blog_item.go
│       ├── create.go
│       ├── delete.go
│       ├── list.go
│       ├── main.go
│       ├── read.go
│       └── update.go
├── go.mod
```

- **proto**: Contains the Protocol Buffer (protobuf) definition for the blog service.
- **server**: Contains the server-side code for the blog service.
- **client**: Contains the client-side code for interacting with the blog service.
- **README.md**: This documentation file.
- **Makefile**: This file is the real deal. It builds the go code, runs the client and the server.

## Protocol Buffer Definition

The `proto/blog.proto` file defines the service and messages used in the project. It specifies the following methods:

- `CreateBlog`: Creates a new blog post.
- `ReadBlog`: Reads a blog post by ID.
- `UpdateBlog`: Updates an existing blog post.
- `DeleteBlog`: Deletes a blog post by ID.
- `ListBlogs`: Lists all blog posts.

## Server Implementation

The server-side code is located in the `server` directory. Here's an overview of the server files:

- `main.go`: The main server application entry point.
- `blog_service.go`: Implements the `BlogServiceServer` interface, providing the logic for the CRUD operations.

## Client Implementation

The client-side code is located in the `client` directory. The client interacts with the server using gRPC. The client code provides functions to call the server's CRUD methods.

## API Endpoints

The server exposes the following gRPC API endpoints:

- `CreateBlog`: Creates a new blog post.
- `ReadBlog`: Reads a blog post by ID.
- `UpdateBlog`: Updates an existing blog post.
- `DeleteBlog`: Deletes a blog post by ID.
- `ListBlogs`: Lists all blog posts.

## Run Locally

Clone the project

```bash
  git@github.com:viralparmarme/go-grpc-mongodb-blog.git
```

Go to the project directory

```bash
  cd go-grpc-mongodb-blog
```

Install dependencies for your backend server

```bash
  go get go.mongodb.org/mongo-driver
```
```bash
  go get google.golang.org/grpc
```
```bash
  go get google.golang.org/protobuf
```

Install [Docker](https://www.docker.com) and ensure that this command runs the MongoDB instance. Please disable any other MongoDB instance running in your system (trust me, I wasted an embarrasing number of hours trying to figure out why my Go server isn't able to establish a connection to the container running MongoDB)

```bash
  docker-compose up
```

Build the code (ensure that make works in your system)

```bash
  make blog
```

Start the server in one terminal window

```bash
  ./bin/blog/server
```

Start the client in a separate window

```bash
  ./bin/blog/client
```

## Used By

This project is used by the following people:

- [Viral Parmar]()

Feel free to raise a [pull request](https://github.com/viralparmarme/go-grpc-mongodb-blog/compare) and adding your name to be visible here for the whole world to see!
## Acknowledgements

 - [Go](https://go.dev)
 - [MongoDB](https://www.mongodb.com)
 - [gRPC](https://grpc.io)


## Contributing

Raise a [pull request](https://github.com/viralparmarme/go-grpc-mongodb-blog/compare) or an [issue](https://github.com/viralparmarme/go-grpc-mongodb-blog/issues/new/choose) to make this project better for everyone...
## Support

For support, email viralparmarme@gmail.com or send me a DM on [Twitter](http://twitter.com/viralparmarme).

