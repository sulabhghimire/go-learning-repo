package main

import (
	"context"
	mainapipb "grpcClient/proto/gen"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	cert := "cert.pem"

	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalln("Error loading TLS Certification:", err)
	}

	conn, err := grpc.NewClient("127.0.0.1:50001", grpc.WithTransportCredentials(creds))
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
