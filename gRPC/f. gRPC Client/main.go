package main

import (
	"context"
	mainapipb "grpcClient/proto/gen"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Did not connect:", err)
	}
	defer conn.Close()

	client := mainapipb.NewCalculateClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &mainapipb.AddRequest{A: 10, B: 12}

	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln("Could not add", err)
	}

	log.Println("Sum : ", res.Sum)
}
