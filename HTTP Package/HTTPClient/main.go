package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// Create a new HTTP Client
	client := &http.Client{}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/asasd")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	//	fmt.Println(resp.StatusCode)

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(body))

}
