package main

import (
	"context"
	"log"
	"net"

	pb "simplegPRCServer/proto/gen"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{Sum: req.A + req.B}, nil
}

func main() {

	serveAddress := "localhost:50001"

	list, err := net.Listen("tcp", serveAddress)
	if err != nil {
		log.Fatal("Failed to listen.", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCalculateServer(grpcServer, &server{})

	// skipping something

	log.Println("Server is running at", serveAddress)
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal("Failed to server.", err)
	}

}
