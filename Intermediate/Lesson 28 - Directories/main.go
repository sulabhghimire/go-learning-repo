package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getWorkingDirectory() {
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println("Current working directly", dir)
}

func main() {

	checkError(os.RemoveAll("subdir1"))
	checkError(os.RemoveAll("subdir"))

	err := os.Mkdir("subdir1", 0755)
	checkError(err)

	checkError(os.WriteFile("subdir1/file", []byte(""), 0755)) // Need to have subdir1 if we want to create inside subdir

	checkError(os.MkdirAll("subdir/parent/child", 0755))
	checkError(os.MkdirAll("subdir/parent/child1", 0755))
	checkError(os.MkdirAll("subdir/parent/child2", 0755))
	checkError(os.MkdirAll("subdir/parent/child3", 0755))

	checkError(os.WriteFile("subdir/parent/file", []byte(""), 0755))
	checkError(os.WriteFile("subdir/parent/child/file", []byte(""), 0755))

	getWorkingDirectory()
	res, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, entry := range res {
		fmt.Println("Entry:", entry.Name(), "Is Dir:", entry.IsDir(), "Type:", entry.Type())
	}

	fmt.Println("------------------------")
	getWorkingDirectory()

	// Change Dir
	checkError(os.Chdir("subdir/parent/child"))
	res, err = os.ReadDir(".")
	checkError(err)

	getWorkingDirectory()

	for _, entry := range res {
		fmt.Println("Entry:", entry.Name(), "Is Dir:", entry.IsDir(), "Type:", entry.Type())
	}

	fmt.Println("------------------------")

	checkError(os.Chdir("../../.."))
	getWorkingDirectory()

	// Walking directories with filepath
	// (filepath.Walk uses filepath.Info more details but less efficient) and (filepath.WalkDir uses os.DirEntry more efficient)
	fmt.Println("Walking Directory")

	pathfile := "subdir"
	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		checkError(err)
		fmt.Println(path)
		return nil
	})
	checkError(err)

}
