package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	main_pb "grpc_server_consumer_clinet/proto/gen"
)

func main() {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := main_pb.NewCalculatorClient(conn)

	// Server side streaming
	ctx := context.Background()
	req := &main_pb.FibonacciRequest{
		Count: 10,
	}

	stream, err := client.GenerateFibonacci(ctx, req)
	if err != nil {
		log.Fatalln("Error calling GenerateFibonacci func", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of stream")
			break
		}
		if err != nil {
			log.Fatalln("Error receiving GenerateFibonacci func", err)
		}
		log.Println(resp.GetNumber())
	}

	// Client side streaming
	stream1, err := client.SendNumbers(ctx)
	if err != nil {
		log.Fatalln("Error creating stream:", err)
	}

	for num := range 9 {
		err = stream1.Send(&main_pb.NumberRequest{Number: int32(num)})
		if err != nil {
			log.Fatalln("Error sending number to stream:", err)
		}
		time.Sleep(time.Second)
	}

	res, err := stream1.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error receiving response:", err)
	}
	log.Println("Sum:", res.GetSum())
}
