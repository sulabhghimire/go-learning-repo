package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool
type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (s *personSorter) Len() int {
	return len(s.people)
}
func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}
func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}
func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

// Redundant is need to swap by other fields in struct
// type ByAge []Person
// type ByName []Person

// func (p ByAge) Len() int {
// 	return len(p)
// }
// func (p ByName) Len() int {
// 	return len(p)
// }

// func (p ByAge) Less(i, j int) bool {
// 	return p[i].Age < p[j].Age
// }
// func (p ByName) Less(i, j int) bool {
// 	return p[i].Name < p[j].Name
// }

// func (p ByAge) Swap(i, j int) {
// 	p[i], p[j] = p[j], p[i]
// }
// func (p ByName) Swap(i, j int) {
// 	p[i], p[j] = p[j], p[i]
// }

func main() {

	// numbers := []int{5, 3, 1, 4, 2}
	// sort.Ints(numbers)
	// fmt.Println("Sorted numbers:", numbers)

	// stringSlice := []string{"John", "Anthony", "Victor", "Walter", "Steve"}
	// sort.Strings(stringSlice)
	// fmt.Println("Sorted strings:", stringSlice)

	// ===========================================================================
	// Sorting by functions

	// --> sort.Interface
	// Contains three methods a) Len() int b) Less(i, j int) bool c) Swap(i, j int)
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 23},
		{Name: "Anna", Age: 29},
	}
	// sort.Sort(ByAge(people))
	// fmt.Println("Sorted by age:", people)
	// sort.Sort(ByName(people))
	// fmt.Println("Sorted by name:", people)

	// Removing the redundancy
	byAge := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	By(byAge).Sort(people)
	fmt.Println("Sorted by age by personSorter:", people)
	byName := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	By(byName).Sort(people)
	fmt.Println("Sorted by name by personSorter:", people)

	// sorting by functions---> sort.Slice
	fruits := []string{"banana", "cherry", "grapes", "guava", "apple"}
	sort.Slice(fruits, func(i, j int) bool {
		return fruits[i][len(fruits[i])-1] < fruits[j][len(fruits[j])-1]
	})
	fmt.Println("Sorted fruits:", fruits)

}
