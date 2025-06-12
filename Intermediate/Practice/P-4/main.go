package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	// scanner.Split(bufio.ScanWords) // Scans one word at a time
	// scanner.Split(bufio.ScanBytes) // Scans one bytes at a time
	scanner.Split(bufio.ScanRunes)

	file, err := os.OpenFile("target.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening/creatin file:", err)
	}

	for scanner.Scan() {
		word := scanner.Text()
		if word == "exit" {
			break
		}

		_, err = file.WriteString(word)
		if err != nil {
			log.Fatalln("Error writing word to file:", err)
			break
		}
	}
	err = scanner.Err()
	if err != nil {
		log.Fatalln("Error scanning word:", err)
	}

}
