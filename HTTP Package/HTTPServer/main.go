package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello Server!")
	})

	const serverAddr string = "127.0.0.1:8080"

	fmt.Println("Server listening on port 3000")
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
