package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	joinedPath := filepath.Join("Documents", "downloads", "file.zip")
	fmt.Println("Joined Path:", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path:", normalizedPath)

	dir, file := filepath.Split("C:\\User\\Documents\\Personal\\file.txt")
	fmt.Println("Directory:", dir)
	fmt.Println("File:", file)
	fmt.Println("Base:", filepath.Base("C:\\User\\Documents\\Personal\\file.txt"))

	relativePath := "./data/file.txt"
	absolutePath := "C:\\User\\Documents\\Personal\\file.txt"

	fmt.Println(filepath.IsAbs(relativePath))
	fmt.Println(filepath.IsAbs(absolutePath))

	newFile := "C:\\User\\Documents\\Personal\\file.txt"
	fmt.Println("Extract extension of ", filepath.Ext(newFile))
	fmt.Println("Trim Extension:", strings.TrimSuffix(newFile, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rel)

	abs, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Abs", abs)

}
