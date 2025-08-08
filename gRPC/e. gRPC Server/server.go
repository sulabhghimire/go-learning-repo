package main

import (
	"context"
	"log"
	"net"

	pb "simplegPRCServer/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{Sum: req.A + req.B}, nil
}

func main() {

	cert := "cert.pem"
	key := "key.pem"

	serveAddress := "127.0.0.1:50001"

	list, err := net.Listen("tcp", serveAddress)
	if err != nil {
		log.Fatal("Failed to listen.", err)
	}

	cred, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatal("Error loading credential.", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(cred))

	pb.RegisterCalculateServer(grpcServer, &server{})

	// skipping something

	log.Println("Server is running at", serveAddress)
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal("Failed to server.", err)
	}

}
