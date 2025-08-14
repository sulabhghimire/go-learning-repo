package main

import (
	"context"
	main_pb "grpcstreams/proto/gen"
	"io"
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

func (s *server) SendNumbers(stream main_pb.Calculator_SendNumbersServer) error {

	var sum int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&main_pb.NumberResponse{Sum: sum})
		}
		if err != nil {
			log.Fatalln(err)
			return err
		}

		log.Println(req.GetNumber())
		sum += req.GetNumber()
	}
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
