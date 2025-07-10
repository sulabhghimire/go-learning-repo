package main

import "fmt"

type Audit struct {
	CreatedBy string `db:"created_by"`
}

type Teacher struct {
	ID        int    `db:"-"` // skip
	FirstName string `db:"first_name,omitempty"`
	LastName  string `db:"last_name,omitempty"`
	Class     int
	RollNo    int `db:"class_roll_no"`
	Marks     int `db:"score,omitempty"`
	Audit         // embedded
}

func main() {

	// t := Teacher{
	// 	FirstName: "Sulabh",
	// 	LastName:  "Ghimire",
	// 	Class:     10,
	// 	Audit:     Audit{CreatedBy: "system"},
	// }

	// res, err := InsertStruct("", &t) // empty tableOverride ⇒ auto
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(res.Query, res.Args)

	u := User{
		ID:     1,
		Name:   "Sulabh Ghimire",
		Email:  "sulabhghimire82@gmail.com",
		Age:    20,
		Income: 0,
		Address: Address{
			city:    "Pokhara",
			Country: "Nepal",
		},
	}

	mapValue := map[string]any{
		"name": "Sulabh",
		"address": map[string]any{
			"city": "Nepal",
		},
	}
	fmt.Println(mapValue)

	fmt.Println(StructToMap(&u))
}
