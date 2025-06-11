package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string  `json:"first_name"`
	Age       int     `json:"age,omitempty"`
	Email     string  `json:"email"`
	Address   Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {

	person := Person{FirstName: "John", Email: "jonh@gmail.com"}

	// Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON", err)
		return
	}
	fmt.Println(string(jsonData))

	person1 := Person{FirstName: "Sulabh", Age: 25, Email: "sulabh@random.com", Address: Address{City: "KTM", State: "Bagmati"}}
	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON", err)
		return
	}
	fmt.Println(string(jsonData1))

	// Unmarshalling
	var employee Employee
	jsonData2 := `{"full_name": "Sulabh Ghimire", "emp_id": "008", "age":30, "address": {"city": "San Jose", "state":"CA"}}`
	err = json.Unmarshal([]byte(jsonData2), &employee)
	if err != nil {
		fmt.Println("Error marshalling to JSON", err)
		return
	}
	fmt.Println(employee)

	// Marshalling List
	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
	}
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		fmt.Println("Error marshalling to JSON", err)
		return
	}
	fmt.Println(string(jsonList))

	// Handaling unknown JSON Structures
	jsonData3 := `{"name": "John", "age":20, "address": {"city": "New York", "state":"NY"}}`
	var data map[string]any
	err = json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		fmt.Println("Error marshalling to JSON", err)
		return
	}
	fmt.Println(data)

}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
