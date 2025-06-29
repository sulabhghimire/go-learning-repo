package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling incoming orders.")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling users.")
	})

	serveAddr := "127.0.0.1:3000"

	// Load the TLS certificates

	/*
		Contains the public key and the certificate
		Certificate includes information about the key, the owner and the issuer of the certificate
		Used by the client to encrypt data to be sent to the server and to verify the server's identity
	*/
	cert := "cert.pem"

	/*
		Contains the private key that is used to decrypt data encrypted with corresponding public key and to sign data to prove
		its origin
	*/
	key := "key.pem"

	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Create a custom server
	server := &http.Server{
		Addr:      serveAddr,
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	// Enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on:", serveAddr)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Couldn't start the server:", err)
	}

	// HTTP 1.1 Server without TLS
	// if err := http.ListenAndServe(serveAddr, nil); err != nil {
	// 	log.Fatalln("Couldn't start the server:", err)
	// }

}

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("Received request with HTTP Version:", httpVersion)

	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Received request with TLS Version:", tlsVersion)
	} else {
		fmt.Println("Received request without TLS")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown TLS Version"
	}
}
