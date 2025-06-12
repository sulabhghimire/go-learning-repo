package main

import (
	"io"
	"log"
	"os"
)

func main() {

	srcFile, err := os.OpenFile("src.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening src file:", err)
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile("dest.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln("Error creating dest file:", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatalln("Error copying from src file to dst File:", err)
	}

}
