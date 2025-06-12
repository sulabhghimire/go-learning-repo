package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fileName := flag.String("filename", "target.txt", "Name of the file in which you want your content to be saved.")
	flag.Parse()

	targetFile, err := os.OpenFile(*fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error creating target file:", err)
	}

	fmt.Println("Typewriter")
	fmt.Println("Keep on typing your text here and the contents will be copied to a file. Copying will stop when you type exit and press enter")
	fmt.Println("Please enter your text: ")

	scanner := bufio.NewScanner(os.Stdin)

	lineWritten := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "exit" {
			break
		}

		if line != "" {
			_, err := targetFile.WriteString(line + "\n")
			if err != nil {
				log.Fatalln("Error writing line to destination file:", err)
			}
			lineWritten++
		}

	}

	err = scanner.Err()
	if err != nil {
		log.Fatalln("Error scanning the input:", err)
	}

	if lineWritten == 0 {
		fmt.Println("No content written to file. Deleting the file.")
		targetFile.Close()
		deleteFile(*fileName)
		return
	} else {
		fmt.Printf("✅ Session ended. %d lines written to %s.\n", lineWritten, *fileName)
	}
	defer targetFile.Close()

}

func deleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Println("Error deleting target file:", err)
	}
	log.Println("Deleting target file sucessfull.")
}
