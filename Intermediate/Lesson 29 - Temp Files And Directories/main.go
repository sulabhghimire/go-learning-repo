package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// Create a temp file
	// If not given a dir in CreateTemp it creates that file in os temp folder
	tempFile, err := os.CreateTemp("", "temporaryFile")
	checkError(err)
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())
	fmt.Println("Temp file created:", tempFile.Name())

	// Creating TempDirectory
	tempDir, err := os.MkdirTemp("", "temporaryFolder")
	checkError(err)
	defer os.RemoveAll(tempDir)
	fmt.Println("Temp directory creates:", tempDir)

}
