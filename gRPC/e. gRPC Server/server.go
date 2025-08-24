package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "simplegPRCServer/proto/gen"
	farewellpb "simplegPRCServer/proto/gen/farewell"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	_ "google.golang.org/grpc/encoding/gzip" // compression in server
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedGreeterServer
	farewellpb.UnimplementedAufWiedersehenServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("no metadata recieved.")
	}
	log.Println("Metadata: ", md)
	val, ok := md["authorization"]
	if !ok {
		log.Println("No value with authorization key in metadata.")
	} else {
		log.Println(val)
	}

	// Set response headers
	responseHeaders := metadata.Pairs("test", "test-value", "test2", "test-value2")
	err := grpc.SendHeader(ctx, responseHeaders)
	if err != nil {
		return nil, err
	}

	trailer := metadata.Pairs("test-trailer", "test-trailer-value")
	err = grpc.SetTrailer(ctx, trailer)
	if err != nil {
		return nil, err
	}
	return &pb.AddResponse{Sum: req.A + req.B}, nil
}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello user %s.", req.Name)}, nil
}

func (s *server) BidGoodBye(ctx context.Context, req *farewellpb.GoodByeRequest) (*farewellpb.GoodByeResponse, error) {
	return &farewellpb.GoodByeResponse{Message: fmt.Sprintf("Bye Bye %s.", req.Name)}, nil
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
	pb.RegisterGreeterServer(grpcServer, &server{})
	farewellpb.RegisterAufWiedersehenServer(grpcServer, &server{})

	// skipping something

	log.Println("Server is running at", serveAddress)
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal("Failed to server.", err)
	}

}
