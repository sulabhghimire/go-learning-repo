````markdown
# Go Data Types

This document outlines the fundamental **data types** in the Go programming language, their characteristics, and their usage. Go is a statically typed language, meaning each variable must have a defined type, which determines the size, range of values, and operations that can be performed.

---

## Table of Contents

- [🔢 Basic Data Types](#basic-data-types)
  - [1. Integers](#1-integers)
  - [2. Floating Point Numbers](#2-floating-point-numbers)
  - [3. Complex Numbers](#3-complex-numbers)
  - [4. Boolean](#4-boolean)
  - [5. Strings](#5-strings)
  - [6. Rune](#6-rune)
- [🧱 Composite Data Types](#composite-data-types)
  - [1. Arrays](#1-arrays)
  - [2. Slices](#2-slices)
  - [3. Maps](#3-maps)
  - [4. Structs](#4-structs)
- [Other Notable Go Constructs](#other-notable-go-constructs)
- [Zero Values](#zero-values)

## 🔢 Basic Data Types

### 1. Integers

- Represent whole numbers without fractional components.
- Go supports both signed and unsigned integers of various sizes:
  - Signed: `int`, `int8`, `int16`, `int32`, `int64`
  - Unsigned: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- The default `int` type is platform-dependent (either 32 or 64 bits).

### 2. Floating Point Numbers

- Represent numbers with fractional parts.
- Go supports:
  - `float32`: single precision
  - `float64`: double precision

### 3. Complex Numbers

- Used for scientific and engineering calculations.
- Types available:
  - `complex64`: composed of two `float32` values
  - `complex128`: composed of two `float64` values
- Use the `math/cmplx` package for complex operations such as:
  - `real()`, `imag()`, `conj()`, `abs()`, etc.

### 4. Boolean

- Represents logical values: `true` or `false`.

### 5. Strings

- Represent a sequence of characters (UTF-8 encoded).
- Immutable: once created, cannot be changed.

### 6. Rune

- Represents a single Unicode character using the `int32` type.
- **Difference from Strings**:
  - A string is a sequence of bytes.
  - A rune represents a single Unicode code point.

---

## 🧱 Composite Data Types

### 1. Arrays

- Fixed-size collection of elements of the same type.
- Example: `var a [5]int` declares an array of five integers.

### 2. Slices

- Dynamic and flexible sequences built on top of arrays.
- Can grow or shrink in size.
- Example: `[]int{1, 2, 3}`

### 3. Maps

- Unordered collection of key-value pairs.
- All keys must be of the same type, and all values must be of the same type.
- Example: `map[string]int{"age": 30}`

### 4. Structs

- User-defined composite types that group together variables of different types under one name.
- Example:
  ```go
  type Person struct {
      Name string
      Age  int
  }
  ```
````

## Other Notable Go Constructs

- Functions
- Pointers
- Channels
- JSON
- Text Templates
- HTML Templates

## Zero Values

1.  Variables declared without an explicit initialization are assigned a default zero value based on their type:
    a. Numeric types default to `0`.
    b. Boolean type defaults to `false`.
    c. String type defaults to an empty string (`""`).
    d. Pointers, slices, functions, maps, and structs: their zero value is `nil`.
