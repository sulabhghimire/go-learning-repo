package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "grpc_gateway_project/proto/gen"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip" // compression in server
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	if err := req.Validate(); err != nil {
		log.Printf("validation failed: %v", err)

		return nil, status.Errorf(codes.InvalidArgument, "Invalid request: %v", err)
	}

	return &pb.HelloResponse{Message: fmt.Sprintf("Hello user %s.", req.Name)}, nil
}

func runGRPCServer(certFile, keyFile string) {

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalln("failed load TLS Certificates:", err)
	}

	serveAddress := "127.0.0.1:50051"

	list, err := net.Listen("tcp", serveAddress)
	if err != nil {
		log.Fatal("Failed to listen.", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterGreeterServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	log.Println("gRPC Server is running at", serveAddress)
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal("Failed to server.", err)
	}

}

func loadTLSCredentials(certFile, keyFile string) tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalln("Failed to load certificated:", err)
	}
	return cert
}

func runGatewayServer(certFile, keyFile string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		// For HTTP
		// grpc.WithTransportCredentials(insecure.NewCredentials()),
		// For HTTPS
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalln("Failed to register gRPC-Gateway handler:", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{loadTLSCredentials(certFile, keyFile)},
	}

	server := &http.Server{
		Addr:      ":8080",
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	log.Printf("HTTPS Server is running on port: 8080")
	// Without tls
	// err = http.ListenAndServe(":8080", mux)
	err = server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatalln("Failed to start REST API server:", err)
	}

}

func main() {

	cert := "cert.pem"
	key := "key.pem"

	go runGRPCServer(cert, key)
	runGatewayServer(cert, key)
}
