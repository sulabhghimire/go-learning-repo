# Conditional Statements

Essential to control the flow of program execution based on certain conditions. Go supports `if`, `if...else`, `if...else if...else`, and `switch`.

### ✅ If

Executes a block of code if the condition is true.

Example:  
if condition {  
  // Code to execute  
}

---

### ✅ If...Else

Follows an if statement and executes a block of code if the if condition is false.

Example:  
if condition {  
  // Code to execute  
} else {  
  // Code to execute  
}

---

### ✅ If...Else If...Else

Evaluates multiple conditions in sequence and executes the first matching block.

Example:  
if condition {  
  // Code to execute  
} else if condition {  
  // Code to execute  
} else {  
  // Code to execute  
}

---

### ✅ Switch Case

Provides a concise way to evaluate multiple possible conditions against a single expression.

- Simplifies multiple `if` and `else if` blocks.
- Evaluates the expression and compares it against each `case` value.
- Executes the first matching `case`, or `default` if no match.
- Unlike other languages, Go automatically breaks after a case block.
- To fall through to the next case, `fallthrough` must be used explicitly.
- `switch` can also be used with type assertions using `switch x.(type)`, but `fallthrough` is not allowed in type switches.

Example:  
switch expression {  
  case value1:  
    // Code to execute  
  case value2:  
    // Code to execute  
  default:  
    // Code to execute  
}

---
