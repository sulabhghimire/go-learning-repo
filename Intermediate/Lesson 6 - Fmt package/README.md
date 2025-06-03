# fmt Package in Go

The `fmt` package in Go is a fundamental standard library package that provides formatted input and output capabilities. It is commonly used for printing to standard output, returning formatted strings, and scanning input.

## Overview

- The `fmt` package is essential for working with input/output in Go.
- It supports printing to standard output (console), formatting strings, and reading user input.
- It provides a consistent and powerful set of tools for formatting text.

---

## Printing Functions

These functions print output to standard output.

- **Print()** – Prints values in default format without a newline.
- **Println()** – Prints values followed by a newline.
- **Printf()** – Prints formatted output using format specifiers (like `%v`, `%d`, `%s`, etc.).

---

## Formatting Functions

These functions return formatted output as strings instead of printing.

- **Sprint()** – Returns a formatted string in default format.
- **Sprintf()** – Returns a string using a format specifier.
- **Sprintln()** – Returns a string with formatted output followed by a newline.

---

## Scanning Functions

These functions read input from standard input (usually the keyboard) and store it into variables.

- **Scan()** – Reads space-separated values into provided variables.
- **Scanf()** – Reads formatted input based on format specifiers.
- **Scanln()** – Similar to `Scan()` but stops reading at newline.

---

## Error Formatting

- **Errorf()** – Formats error messages using format specifiers. Returns an error of type `error`.

---

## Summary

The `fmt` package is a core utility in Go development. It simplifies common tasks like printing messages, reading user input, formatting strings, and generating error messages. Familiarity with `fmt` is essential for effective Go programming.
