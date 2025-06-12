package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("file.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}
	defer file.Close()

	bufScanner := bufio.NewScanner(file)

	for bufScanner.Scan() {

		line := bufScanner.Text()
		fmt.Println(line)

	}

	err = bufScanner.Err()
	if err != nil {
		log.Fatalln("Error scanning file:", err)
	}
}
