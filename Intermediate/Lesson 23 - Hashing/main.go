package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	password := "password123"

	// hash256 := sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))

	// fmt.Println("Password:", password, "256 HASH:", hash256, "512 HASH:", hash512)
	// fmt.Printf("HEX Value 256: %X\nHEX Value 512: %X\n", hash256, hash512)

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// hash the passwrod with salt
	signUpHash := generateHash(password, salt)

	// Store the salt and password in database
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)    // Simulate as store in database
	fmt.Println("Hash:", signUpHash) // Simulate as store in database

	// Verify the password
	// retrive the saltStr and decode it
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return
	}

	logInHash := generateHash(password, decodedSalt)

	// Compre the stored signUpHash with the generated logInHash
	if signUpHash == logInHash {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Login failed. Please check user credentials.")
	}

}

// Function to generate salt
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	// Read full read cryptographic random number from rand.Reader and save this into salt and reads the size of byte of salt
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Function to hash the input
func generateHash(input string, salt []byte) string {

	saltedInput := append(salt, []byte(input)...)
	hash := sha256.Sum256(saltedInput)

	return base64.StdEncoding.EncodeToString(hash[:])

}
