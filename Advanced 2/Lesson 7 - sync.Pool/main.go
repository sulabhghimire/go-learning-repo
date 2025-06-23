package main

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func main() {

	var pool = sync.Pool{
		New: func() any {
			fmt.Println("Creating a new person")
			return &person{}
		},
	}

	// Get an object from the pool
	person1 := pool.Get().(*person)
	person1.name = "John"
	person1.age = 18

	fmt.Println("Got person: one", person1)
	fmt.Printf("Person1 - Name %s, Age %d\n", person1.name, person1.age)

	pool.Put(person1)
	fmt.Println("Returned person1 to pool")

	person2 := pool.Get().(*person)
	fmt.Println("Got person2:", person2)

	person3 := pool.Get().(*person)
	fmt.Println("Got person3:", person3)
	person3.name = "Jane"

	// Returning object to pool again
	pool.Put(person2)
	pool.Put(person3)

	person4 := pool.Get().(*person)
	fmt.Println("Got person 4:", person4)

	person5 := pool.Get().(*person)
	fmt.Println("Got person 5:", person5)
}
