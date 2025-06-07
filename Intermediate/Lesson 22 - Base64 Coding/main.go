package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("He~lo, Base64 encoding")

	// Encode Base64
	encodedStr := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Standard encoded string:", encodedStr)

	// Decode from Base64
	decoded, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		fmt.Println("Error Decoding:", err)
		return
	}
	fmt.Println("Decoded String:", string(decoded))

	// URL Safe encoding
	// Trying to avoid "/" and "+" for URL Safe

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe Encoded:", urlSafeEncoded)

}
