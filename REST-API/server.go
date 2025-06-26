package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling incoming orders.")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling users.")
	})

	serveAddr := "127.0.0.1:3000"

	fmt.Println("Server is running on:", serveAddr)
	if err := http.ListenAndServe(serveAddr, nil); err != nil {
		log.Fatalln("Couldn't start the server:", err)
	}

}
