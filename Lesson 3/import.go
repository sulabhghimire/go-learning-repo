package main

// import "fmt" --> for one packages

// for multiple packages import
import (
	"fmt"
	foo "net/http" // foo is an alias for net/http package i.e a named import
)

func main() {
	fmt.Println("Hello, Go standard libarary!")

	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}
