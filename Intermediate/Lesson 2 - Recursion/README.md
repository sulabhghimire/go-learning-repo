# Recursion in Programming

Recursion is a **fundamental concept** in programming where a function calls itself **directly or indirectly** to solve a problem by breaking it down into **smaller sub-problems** of the same type.

---

## 📖 Definition

> **Recursion** is the process where a function solves a problem by calling itself with simpler inputs, continuing until it reaches a stopping condition called the **base case**.

---

## 🧩 How Recursion Works

1. **Recursive Case**  
   The part of the function where it calls itself with a **simpler or smaller** input, making progress toward the base case.

2. **Base Case**  
   The stopping condition that ends the recursive calls. Without a base case, recursion would continue indefinitely and **cause a stack overflow**.

---

## ✅ Practical Uses

- **Mathematical Algorithms**  
  (e.g., factorial, Fibonacci sequence, power functions)

- **Tree and Graph Algorithms**  
  (e.g., traversals like DFS, parsing hierarchical data)

- **Divide and Conquer Algorithms**  
  (e.g., quicksort, mergesort, binary search)

---

## 🌟 Benefits

- **Simplicity**  
  Naturally expresses problems that are recursive in nature.

- **Clarity**  
  Makes complex logic easier to understand when written clearly.

- **Flexibility**  
  Handles nested or hierarchical data structures effectively.

---

## ⚠️ Considerations

- **Performance**  
  Recursive solutions can be less efficient due to call stack overhead.

- **Base Case**  
  Always ensure a clear and reachable base case to prevent infinite recursion.

---

## 🧠 Best Practices

1. **Testing**  
   Verify that the base case works and the function eventually reaches it.

2. **Optimization**  
   Consider memoization or converting to iteration if performance becomes an issue.

3. **Well-defined Recursive Case**  
   Ensure each recursive step moves closer to the base case.

---

## 📌 Example in Go

```go
package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1 // base case
    }
    return n * factorial(n-1) // recursive case
}

func main() {
    fmt.Println(factorial(5)) // Output: 120
}
```
