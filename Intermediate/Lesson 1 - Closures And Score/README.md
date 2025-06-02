# Closures and Scope in Go

Closures are a powerful feature in Go that allow functions to **capture and manipulate variables** defined outside their immediate function body.

## 📖 Definition

> A **closure** is a function value that references variables from outside its body.

Closures can **assign to** and **access** captured variables. These variables **persist** as long as the closure is referenced, enabling powerful patterns such as stateful functions and encapsulation.

## 🔍 How Closures Work

Closures work based on **lexical scoping**, meaning they capture variables from their **surrounding context**—where the function was **defined**, not where it is called. This allows closures to access variables even **after the outer function has finished execution**.

---

## ✅ Use Cases

1. **Create Stateful Functions**  
   Maintain private state across multiple invocations.

2. **Encapsulate Functionality and Data**  
   Package logic with internal state, avoiding global variables.

3. **Callback Functions**  
   Useful for event handlers, goroutines, or asynchronous execution.

---

## ⚠️ Considerations

- **Memory Usage**  
  Captured variables remain in memory as long as the closure is referenced.

- **Concurrency**  
  Be cautious with shared state in closures when working with goroutines.

---

## 🧠 Best Practices

1. **Limit Scope**  
   Only capture variables that are necessary to avoid unintended dependencies.

2. **Avoid Overuse**  
   Don’t replace simple functions with closures unnecessarily—use them when there's a clear benefit (e.g., state retention or deferred execution).

---

## 📌 Example in Go

```go
package main

import "fmt"

func adder() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    add := adder()
    fmt.Println(add()) // 1
    fmt.Println(add()) // 2
    fmt.Println(add()) // 3
}
```
