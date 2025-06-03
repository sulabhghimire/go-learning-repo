package main

import "fmt"

// declearing a struct
type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address // Embedding Address struct in Person
	PhoneHomeCell         // Anonymous Struct
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}

func main() {

	// Initializing a struct
	p1 := Person{
		firstName: "Sulabh",
		lastName:  "Ghimire",
		age:       25,
		address: Address{ // Embedded struct
			city:    "Pokhara",
			country: "Nepal",
		},
		PhoneHomeCell: PhoneHomeCell{ // Anonymous struct
			home: "123234234",
			cell: "213213324",
		},
	}

	p2 := Person{
		firstName: "John",
		age:       45,
	}
	p2.address.city = "Kathmandu"
	p2.address.country = "Nepal"

	p3 := Person{
		firstName: "John",
		age:       45,
	}
	p3.address.city = "Kathmandu"
	p3.address.country = "Nepal"

	// Accesing the fields of a struct
	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1.address.city)
	fmt.Println(p2.address)
	fmt.Println(p1.cell) // Anonymous struct
	fmt.Println("Are p1 and p2 equal", p1 == p2)
	fmt.Println("Are p2 and p3 equal", p3 == p2)

	// Anonymous Structs
	user := struct {
		userName string
		email    string
	}{
		userName: "user123",
		email:    "pseudoemail",
	}
	fmt.Println(user.userName)

	// Methods
	fmt.Println(p1.fullName())

	// Pointer recievers
	fmt.Println("Before increment", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After increment", p1.age)

}
