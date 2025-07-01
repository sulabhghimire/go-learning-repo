package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/*

	json.Marshal and json.Unmarshal are used to in-memory serialization and deserialization
		-> Good for quick and small sized data sets
	json.NewDecoder and json.NewEncoder are used for streaming json data
		-> ideal for handling large datasets or working with network connections
		-> json.NewDecoder creates a NewDecoder that reads from io.Reader
			-> useful for streaming data such as reading data from network connection or a file

*/

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	user := User{
		Name:  "Alice",
		Email: "alice@example.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(jsonData))

	var user1 User
	err = json.Unmarshal(jsonData, &user1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(user1)

	jsonData1 := `{"name":"John", "email":"john@example.com"}`
	reader := strings.NewReader(jsonData1)

	decoder := json.NewDecoder(reader)

	var user2 User
	err = decoder.Decode(&user2)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(user2)

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(user)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Print(buf.String())

}
