package main

import (
	"context"
	main_pb "grpcstreams/proto/gen"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	main_pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *main_pb.AddRequest) (*main_pb.AddResponse, error) {
	return &main_pb.AddResponse{Sum: req.A + req.B}, nil
}

func (s *server) GenerateFibonacci(req *main_pb.FibonacciRequest, stream main_pb.Calculator_GenerateFibonacciServer) error {
	n := req.Count
	a, b := 0, 1

	for range int(n) {
		err := stream.Send(&main_pb.FibonacciResponse{Number: int32(a)})
		if err != nil {
			return nil
		}
		a, b = b, a+b
		time.Sleep(time.Second) // Simulate processing time
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	main_pb.RegisterCalculatorServer(grpcServer, &server{})

	log.Println("Running gPRC server at:", "localhost:50051")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}

}
