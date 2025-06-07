package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Initiate A NEW Reader

	// strings.NewReader is a function from strings package that takes a strings and returns a new Reader object
	// bufio.NewReader accepts other reader as input and it buffers that input
	// Source will be external file, image or anything
	reader := bufio.NewReader(strings.NewReader("Hello, bufio packageee!\n")) // String is source

	// Read the data byte slice
	data := make([]byte, 20)
	n, err := reader.Read(data) // returns number of bytes read,and err
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	line, err := reader.ReadString('\n') // retuns string and error
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Println("Read string:", line)
	// Outputs ! as first reads upto 20 bytes and again when reading starts it starts from  point where reading was left

	fmt.Println()
	fmt.Println("-------------------------------------------------------")
	fmt.Println()
	// bufio.Writer

	writer := bufio.NewWriter(os.Stdout)

	// Writing byte slice first
	writeData := []byte("Hello, bufio package.\n")
	n, err = writer.Write(writeData) // returns number of bytes writeen and error
	if err != nil {
		fmt.Println("Error writing bytes:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// All the data written to writer is written to internal buffer and it is not immediately wirtten to os.Stdout
	// So we need to flush the buffer data to underlying writer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error Flushing Writer:", err)
		return
	}

	// Writing strings
	str := "This is a string.\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error Flushing Writer:", err)
		return
	}
}
