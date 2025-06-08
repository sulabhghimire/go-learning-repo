package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Writing data to the file
	data := []byte("Hello World!\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing bytes to file:", err)
		return
	}

	fmt.Println("Data has been written to file with write")

	file1, err := os.Create("write_string.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file1.Close()

	_, err = file1.WriteString("Hello Go!\n")
	if err != nil {
		fmt.Println("Error writing string to file:", err)
		return
	}

	fmt.Println("Data has been written to file with WriteString")

}
