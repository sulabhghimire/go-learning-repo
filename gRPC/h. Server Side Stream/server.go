package main

import (
	"bufio"
	"context"
	"fmt"
	main_pb "grpcstreams/proto/gen"
	"io"
	"log"
	"net"
	"os"
	"strings"
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

func (s *server) Chat(stream main_pb.Calculator_ChatServer) error {

	reader := bufio.NewReader(os.Stdin)

	for {
		// Receiving values/messages from the stream
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
			return err
		}
		log.Println("Received Message:", req.GetMessage())

		// Read input from the terminal
		fmt.Println("Enter Response:")
		inputMsg, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Error reading message from user in server:", err)
			return err
		}

		responseMsg := &main_pb.ChatMessage{Message: strings.TrimSpace(inputMsg)}

		// Sending data/values through the stream
		err = stream.Send(responseMsg)
		if err != nil {
			log.Fatalln("Error sending value to stream:", err)
			return err
		}

	}

	fmt.Println("Returning control")
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
