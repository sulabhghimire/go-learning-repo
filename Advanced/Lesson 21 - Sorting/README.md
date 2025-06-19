# Custom Sorting in Go: From Boilerplate to Reusability

This repository demonstrates different ways to sort custom data structures in Go, focusing on the evolution from a repetitive, boilerplate-heavy approach to a flexible, functional, and reusable pattern. It uses the standard library's `sort` package.

## The Scenario

We have a slice of `Person` structs, and we want to sort this slice based on different fields, such as `Age` and `Name`.

```go
type Person struct {
	Name string
	Age  int
}

var people = []Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 23},
    {Name: "Anna", Age: 29},
}
```

To use Go's `sort.Sort()` function, we need a type that satisfies the `sort.Interface`:

```go
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
```

## Method 1: The Redundant (Boilerplate) Approach

A common but verbose way to implement custom sorting is to create a new type for _each field_ you want to sort by. Each of these new types must then implement all three methods of `sort.Interface`.

### Implementation

To sort by age, we create `ByAge`. To sort by name, we create `ByName`.

```go
// To sort by Age
type ByAge []Person

func (p ByAge) Len() int           { return len(p) }
func (p ByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByAge) Less(i, j int) bool { return p[i].Age < p[j].Age }

// To sort by Name
type ByName []Person

func (p ByName) Len() int           { return len(p) }
func (p ByName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByName) Less(i, j int) bool { return p[i].Name < p[j].Name }
```

### Usage

```go
// Sort by age
sort.Sort(ByAge(people))
fmt.Println("Sorted by age:", people)

// Sort by name
sort.Sort(ByName(people))
fmt.Println("Sorted by name:", people)
```

### The Problem: Redundancy

Notice the repetition. The `Len()` and `Swap()` methods are **identical** for both `ByAge` and `ByName`. The only thing that changes is the logic inside the `Less()` method. If we wanted to add sorting by a third or fourth field, we would have to copy and paste `Len()` and `Swap()` again, leading to fragile and hard-to-maintain code. This violates the **DRY (Don't Repeat Yourself)** principle.

---

## Method 2: The Functional Approach (Reusable and Clean)

A much better approach is to decouple the sorting logic from the sorting mechanism. We can create a single, generic `sorter` struct that implements `sort.Interface` once. This `sorter` will hold a function that defines the comparison logic, allowing us to pass in any sorting criteria we want.

### Implementation

1.  **Define a function type `By`** that represents the signature of our comparison logic. This makes the code more readable and self-documenting.
2.  **Create a `personSorter` struct** that holds the slice of people and an instance of our `By` function.
3.  **Implement `sort.Interface` on `personSorter` just once.** The `Less` method simply calls the comparison function it holds.

```go
// By is a function type that defines the sort order.
type By func(p1, p2 *Person) bool

// personSorter joins a slice of People to a By function.
type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool // or just `by By`
}

// Len, Swap, and Less implement the sort.Interface.
func (s *personSorter) Len() int {
	return len(s.people)
}
func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}
func (s *personSorter) Less(i, j int) bool {
    // The key part: the comparison logic is delegated to the `by` function!
	return s.by(&s.people[i], &s.people[j])
}
```

We also add a convenience `Sort` method to our `By` function type to tie everything together.

```go
// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}
```

### Usage

Now, instead of defining new types, we just define small functions (or closures) for our comparison logic and pass them to our universal sorter.

```go
// Define the logic for sorting by age
byAge := func(p1, p2 *Person) bool {
    return p1.Age < p2.Age
}
// Use the sorter with the byAge logic
By(byAge).Sort(people)
fmt.Println("Sorted by age:", people)

// Define the logic for sorting by name
byName := func(p1, p2 *Person) bool {
    return p1.Name < p2.Name
}
// Use the same sorter with the byName logic
By(byName).Sort(people)
fmt.Println("Sorted by name:", people)
```

### The Benefit: Flexibility and No Redundancy

This pattern is vastly superior:

- **DRY:** The `Len` and `Swap` logic is written only once.
- **Flexible:** To add a new sort order, you only need to write a new comparison function. No new types or `sort.Interface` implementations are needed.
- **Clear:** The intent is very clear at the call site. You can see exactly what logic is being used to perform the sort.

## Modern Alternative: `sort.Slice`

Since Go 1.8, the `sort` package includes `sort.Slice`, which provides a convenient way to sort slices using only a comparison function. This often eliminates the need to create a custom sorter struct at all. It uses reflection internally but is highly optimized and is the idiomatic choice for many simple sorting tasks.

### Usage

`sort.Slice` takes the slice to be sorted and a `less` function that compares elements at indices `i` and `j`.

```go
fruits := []string{"banana", "cherry", "grapes", "guava", "apple"}

// Sort by the last letter of the fruit name
sort.Slice(fruits, func(i, j int) bool {
    // The less function has direct access to the slice
    return fruits[i][len(fruits[i])-1] < fruits[j][len(fruits[j])-1]
})
fmt.Println("Sorted fruits:", fruits) // [banana, apple, cherry, grapes, guava]
```

For our `Person` struct, it would look like this:

```go
// Sorting people by age using sort.Slice
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
```

This is even more concise than the functional `personSorter` pattern, making it ideal for one-off or simple sorting needs. The `personSorter` pattern remains valuable for more complex scenarios or when you want to bundle sorting logic into a reusable package.

## How to Run

To run the example code and see the output:

```sh
go run main.go
```
