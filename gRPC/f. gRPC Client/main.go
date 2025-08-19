package main

import (
	"context"
	mainapipb "grpcClient/proto/gen"
	farewellpb "grpcClient/proto/gen/farewell"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
)

func main() {

	cert := "cert.pem"

	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalln("Error loading TLS Certification:", err)
	}

	conn, err := grpc.NewClient("127.0.0.1:50001", grpc.WithTransportCredentials(creds), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	if err != nil {
		log.Fatalln("Did not connect:", err)
	}
	defer conn.Close()

	client := mainapipb.NewCalculateClient(conn)
	client2 := mainapipb.NewGreeterClient(conn)
	client3 := farewellpb.NewAufWiedersehenClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &mainapipb.AddRequest{A: 10, B: 12}

	res, err := client.Add(ctx, req, grpc.UseCompressor(gzip.Name))
	if err != nil {
		log.Fatalln("Could not add", err)
	}
	log.Println("Sum : ", res.Sum)

	req2 := &mainapipb.HelloRequest{
		Name: "John",
	}
	res2, err := client2.Greet(ctx, req2)
	if err != nil {
		log.Fatalln("Could not greet", err)
	}
	log.Println("Greeting message : ", res2.Message)

	req3 := &farewellpb.GoodByeRequest{
		Name: "Sulabh",
	}
	res3, err := client3.BidGoodBye(ctx, req3)
	if err != nil {
		log.Fatalln("Could not bid goodbye", err)
	}
	log.Println("Goodbye message : ", res3.Message)

}
