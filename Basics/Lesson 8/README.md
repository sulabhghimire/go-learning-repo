# Operators in Go (Golang)

This document provides a quick reference to **Logical**, **Bitwise**, and **Comparison** operators in the Go programming language.

---

## 🔹 Logical Operators

Logical operators in Go are used with boolean values.

| Operator | Name | Description                | Example                  |
| -------- | ---- | -------------------------- | ------------------------ | ---------------------------- | ----- | --- | -------------- |
| `!`      | NOT  | Inverts the boolean value  | `!true // false`         |
| `        |      | `                          | OR                       | True if at least one is true | `true |     | false // true` |
| `&&`     | AND  | True only if both are true | `true && false // false` |

---

## 🔹 Bitwise Operators

Bitwise operators work on binary representations of integers.

| Operator | Name        | Description                                       | Example                                |
| -------- | ----------- | ------------------------------------------------- | -------------------------------------- | --- | ------- |
| `&`      | AND         | Each bit is 1 if both bits are 1                  | `5 & 3 // 1`                           |
| `        | `           | OR                                                | Each bit is 1 if at least one bit is 1 | `5  | 3 // 7` |
| `^`      | XOR         | Each bit is 1 if only one of the bits is 1        | `5 ^ 3 // 6`                           |
| `&^`     | AND NOT     | Clears bits where the second operand has bits set | `5 &^ 3 // 4`                          |
| `<<`     | Left Shift  | Shifts bits to the left                           | `5 << 1 // 10`                         |
| `>>`     | Right Shift | Shifts bits to the right                          | `5 >> 1 // 2`                          |

> **Note:** Go uses `&^` as the "bit clear" or "AND NOT" operator, which is unique compared to many other languages.

---

## 🔹 Comparison Operators

These operators return a boolean result (`true` or `false`).

| Operator | Name                  | Example  | Result  |
| -------- | --------------------- | -------- | ------- |
| `==`     | Equal                 | `5 == 5` | `true`  |
| `!=`     | Not Equal             | `5 != 3` | `true`  |
| `<`      | Less Than             | `3 < 5`  | `true`  |
| `<=`     | Less Than or Equal    | `5 <= 5` | `true`  |
| `>`      | Greater Than          | `6 > 2`  | `true`  |
| `>=`     | Greater Than or Equal | `7 >= 8` | `false` |

---

## 🧪 Example Code

```go
package main

import (
	"fmt"
)

func main() {
	a, b := 5, 3

	// Logical
	fmt.Println("Logical NOT:", !true)
	fmt.Println("Logical OR:", true || false)
	fmt.Println("Logical AND:", true && false)

	// Bitwise
	fmt.Println("Bitwise AND:", a&b)   // 1
	fmt.Println("Bitwise OR:", a|b)    // 7
	fmt.Println("Bitwise XOR:", a^b)   // 6
	fmt.Println("Bit Clear (&^):", a&^b) // 4
	fmt.Println("Left Shift:", a<<1)   // 10
	fmt.Println("Right Shift:", a>>1)  // 2

	// Comparison
	fmt.Println("Equal:", a == b)      // false
	fmt.Println("Not Equal:", a != b)  // true
	fmt.Println("Less Than:", a < b)   // false
	fmt.Println("Greater Than:", a > b) // true
}
```
