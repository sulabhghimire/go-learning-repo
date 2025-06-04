package main

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element, true
}

func (s Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("The stack is empty.")
		return
	}
	fmt.Print("Stack Elements: ")
	for _, element := range s.elements {
		fmt.Print(element, " ")
	}
	fmt.Println()
}

func main() {

	// x, y := 1, 2
	// x, y = swap(x, y)
	// fmt.Println("X:", x, "Y:", y)

	// x1, y1 := "A", "B"
	// x1, y1 = swap(x1, y1)
	// fmt.Println("X1:", x1, "Y1:", y1)

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printAll()
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty", intStack.isEmpty())
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty", intStack.isEmpty())

}

/*
Benifits of Generics
a. Code Reusability
b. Type Safety
c. Peformance

Considerations
a. Type Constraints
b. Documentation
c. Testing
*/
