# For Loop in Go

## 1. What is a For Loop?

A **for loop** is a control structure that allows us to execute a block of code repeatedly based on a condition.

---

## 2. Syntax

for initialization; condition; post {
// Code block to execute repeatedly
}

- **Initialization**: Executed once before the loop starts. Typically used to initialize loop variables.
- **Condition**: Checked before each iteration. The loop continues as long as this condition evaluates to true.
- **Post**: Executed after each iteration. Usually increments or updates the loop variables.

### Example

for i := 1; i <= 5; i++ {
fmt.Println(i)
}

### Control Statements

- **break**: Terminates the loop immediately and transfers control to the first statement after the loop.
- **continue**: Skips the rest of the current iteration and moves to the next iteration.

---

## 3. Using For Loop as a While Loop

In Go, `for` can also be used like a `while` loop by omitting the initialization and post statements:

i := 1
for i <= 5 {
fmt.Println(i)
i++
}

---

## 4. Infinite Loop

If the condition is also omitted, the loop runs indefinitely:

for {
// Infinite loop
}

Use `break` inside the loop to exit when a condition is met.

---
