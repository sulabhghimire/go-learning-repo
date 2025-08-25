package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "simplegPRCServer/proto/gen"

	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/encoding/gzip" // compression in server
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	if err := req.Validate(); err != nil {
		log.Printf("validation failed: %v", err)

		return nil, status.Errorf(codes.InvalidArgument, "Invalid request: %v", err)
	}

	return &pb.HelloResponse{Message: fmt.Sprintf("Hello user %s.", req.Name)}, nil
}

func main() {

	serveAddress := "127.0.0.1:50001"

	list, err := net.Listen("tcp", serveAddress)
	if err != nil {
		log.Fatal("Failed to listen.", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreeterServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	log.Println("Server is running at", serveAddress)
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal("Failed to server.", err)
	}

}
