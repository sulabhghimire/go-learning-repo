package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File Opened sucessfully.")

	// // Reading the contents of the opened files
	// data := make([]byte, 124) // Buffer to read data into
	// _, err = file.Read(data)

	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }
	// fmt.Println("File content:", string(data))

	// Reading the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

}
