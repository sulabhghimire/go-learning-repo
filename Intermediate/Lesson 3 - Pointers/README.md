# Pointers in Go

## 1. What is a Pointer?

A pointer is a variable that stores the memory address of another variable.

## 2. Pointer Types

A pointer has a specific type. It can be:

- An integer pointer
- A floating point pointer
- A string pointer
- And so on...

## 3. Use Cases

Pointers are commonly used to:

- Modify the value of a variable indirectly
- Pass large data structures efficiently between functions
- Manage memory directly for performance reasons

## 4. Pointer Declaration

### a. Declaration Syntax

```go
var ptr *variableType

```

`ptr` is a pointer to an integer (\*int)

## 5. Address and Dereferencing

- Use `&` to get the address of a variable.
- Use `*` to dereference a pointer.

## 6. Pointer Limitations in Go

Go does not support pointer arithmetic. Pointers are limited to referencing and dereferencing only.

## 7. Passing Pointers to Functions

Pointers can be passed to functions, allowing them to modify the original variable.

## 8. Unsafe Pointers

unsafe.Pointer(&x) converts the address of x to an unsafe.Pointer.
