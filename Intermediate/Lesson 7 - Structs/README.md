# Structs in Go

Structs in Go are composite data types that allow grouping variables of different types under a single name. They serve a similar purpose to classes in object-oriented languages but are more lightweight and do not support inheritance.

## Overview

Structs are a powerful feature in Go, providing a clean way to model complex data structures and encapsulate behavior using methods. They're commonly used to define and manage the shape of data in Go programs.

## Features and Concepts

- **Definition**: A struct is defined using the `type` and `struct` keywords. It consists of a list of fields enclosed in curly braces.

- **Fields**: Each field has a name and a type. Field names must be unique within a struct.

- **Initialization**:

  - Structs can be initialized using struct literals.
  - Fields not explicitly initialized will take their zero value (e.g., `0`, `""`, `false`, `nil`).

- **Anonymous Structs**:

  - Go supports defining structs without a named type. These are useful for short-lived or one-off data structures.

- **Methods on Structs**:

  - Functions can be associated with structs using a receiver.
  - Receivers can be value or pointer types.
  - Pointer receivers allow modification of the struct’s fields from within the method.

- **Exported Fields**:

  - Field names starting with a capital letter are exported and accessible from other packages.
  - Field names starting with a lowercase letter are unexported (private).

- **Embedded Structs**:

  - One struct can be embedded inside another, promoting its fields and methods to the outer struct.
  - This promotes composition over inheritance and allows reuse of common fields or behaviors.

- **Anonymous Fields**:

  - Structs can contain anonymous fields (fields with only a type and no explicit name). This simplifies access, as the embedded type’s fields are promoted to the outer struct.

- **Comparability**:
  - Structs can be compared using the `==` operator if all their fields are comparable types.

## Best Practices

- Use structs to model real-world entities and data.
- Prefer composition using embedded structs rather than inheritance.
- Use pointer receivers for methods that modify the struct or for performance when working with large structs.
- Keep struct definitions in a separate file if used across packages to maintain modular code organization.

## Summary

Structs are foundational to data modeling in Go. Their simplicity, combined with method support and composition, provides a clean and efficient way to build scalable and maintainable applications.
