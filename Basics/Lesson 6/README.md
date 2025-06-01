# Naming Conventions in Go

This document outlines common naming conventions used in Go programming to ensure code clarity, consistency, and readability.

---

## 1. Pascal Case (PascalCase)

- Used to name types such as structs, interfaces, and enums.
- Examples include: `CalculateArea`, `UserInfo`, `NewHTTPRequest`.
- Typically applied for defining types or exported identifiers.

Example:  
Type Employee struct {  
 FirstName string  
 LastName string  
 Age int  
}

---

## 2. Snake Case (snake_case)

- Used for naming variables, constants, and filenames.
- Examples include: `user_id`, `first_name`, `http_request`.

---

## 3. UPPERCASE

- Used exclusively for naming constants.
- Uppercase emphasizes immutability and makes constants stand out.
- Examples include: `MAX_RETRIES = 5`, `DEFAULT_TIMEOUT`, `DB_HOST`.

---

## 4. Mixed Case (camelCase)

- Used for variables or identifiers especially when interfacing with external libraries or APIs.
- Examples include: `javaScript`, `htmlDocument`, `isValid`, `employeeID = 1001`.

---

## Important Tips

- Maintain consistency throughout the codebase by following these conventions.
- Minimize the use of abbreviations unless they are widely recognized.
- Keep package names short, concise, and in lowercase without underscores or mixed casing.

---

# Constants in Go

Constants are immutable values in Go that do not change during the execution of a program. They remain consistent throughout the lifecycle of the program.

---

## Key Points

- Constants hold values that are fixed and cannot be modified after they are defined.
- Their values must be known at compile time, meaning you cannot assign values that are determined during runtime.
- Go supports both typed and untyped constants.

---

## Typed Constants

Typed constants have a specific data type assigned to them.  
Example:  
const Pi float64 = 3.14  
const MaxConnections int = 100

---

## Untyped Constants

Untyped constants do not have an explicit type and can be used flexibly in expressions.  
Example:  
const DefaultPort = 8080  
const Greeting = "Hello, World!"

---

## Important Notes

- Constants can be of basic types like boolean, numeric (int, float), and string.
- Once declared, constant values cannot be changed.
- They are useful for defining fixed values

---

# Arithmetic Operators in Go

Arithmetic operators in Go are used to perform mathematical operations on numeric values. This includes both basic operations and considerations like operator precedence and data type behavior.

---

## 1. Basic Operators

- Addition
- Subtraction
- Multiplication
- Division
- Remainder (Modulo)

---

## 2. Operator Precedence

Operations are evaluated based on the following precedence rules:

1. Parentheses
2. Multiplication, Division, Remainder
3. Addition, Subtraction

Example:  
The expression `2 + 3 * 4` evaluates to `14`, not `20`, because multiplication has higher precedence than addition.

To alter the precedence, use parentheses:  
`(2 + 3) * 4` evaluates to `20`.

---

## Important Notes

### 1. Integer Division Truncation

When dividing two integers, the fractional part is discarded.  
Example:  
`const p float64 = 22 / 7`

Here, both `22` and `7` are integers, so the result is truncated to `3`. To preserve precision:

- Use at least one operand as a float:  
  `const p float64 = 22.0 / 7`  
  or  
  `const p float64 = 22 / 7.0`

---

### 2. Overflow and Underflow

Be cautious of numeric limits when performing operations:

#### a. Overflow

Occurs when the result exceeds the maximum value that a numeric data type can hold.  
Example:  
Adding two large integers may wrap the value around to a negative number (in the case of signed integers) or cause undefined behavior (in the case of unsigned integers).

#### b. Underflow

Occurs when the result is smaller than the minimum value storable in the data type.  
This is especially relevant for floating-point numbers and can result in:

- Loss of precision
- Inaccurate calculations
- Zero approximation for extremely small results

---

Understanding and correctly using arithmetic operators ensures reliable and accurate computations in Go programs.
