package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {

	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from the reader:", err)
	}
	fmt.Println("Read data:", string(buf[:n]))

}

func writeToWriter(w io.Writer, data string) {

	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to the reader:", err)
	}

}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error closing the resource:", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer // Creates memory on the stack
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("World")

	mr := io.MultiReader(r1, r2)

	// Since we are reading from multiple resources we need a pointer that is going to store read values
	// into a memory address and we are going to access that memory address and extract value stored in buffers
	buf := new(bytes.Buffer) // allocates memmory on the heap
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading from multi reader:", err)
	}
	fmt.Println(buf.String())
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func writeToFile(filepath, data string) {

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening/creating file:", err)
	}
	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error wwiting to file:", err)
	}

	writer := io.Writer(file)
	writeToWriter(writer, data)
}

func main() {

	fmt.Println("=== Read from Reader ===")
	readFromReader(strings.NewReader("Hello Reader"))

	fmt.Println("=== Write to Writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello Writer")
	fmt.Println(writer.String())

	fmt.Println("=== Buffer Example ===")
	bufferExample()

	fmt.Println("=== Multi Reade Example ===")
	multiReaderExample()

	fmt.Println("=== Pipe Example ===")
	pipeExample()

	fmt.Println("=== Write to file example ===")
	writeToFile("example.txt", "Example data\n")
}
