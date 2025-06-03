# Go Formatting Verbs Reference

This document outlines the formatting verbs used in Go's `fmt` package for formatting different types of values.

---

## General Formatting Verbs

- `%v` &nbsp;– Prints the value in default format
- `%#v` – Prints the value in Go-syntax format
- `%T` &nbsp;– Prints the type of the value
- `%%` &nbsp;– Prints a literal percent sign

---

## Integer Formatting Verbs

- `%b` &nbsp;– Base 2
- `%d` &nbsp;– Base 10
- `%+d` – Base 10 with sign
- `%o` &nbsp;– Base 8
- `%O` &nbsp;– Base 8 with leading `0o`
- `%x` &nbsp;– Base 16 (lowercase)
- `%X` &nbsp;– Base 16 (uppercase)
- `%#x` – Base 16 with leading `0x` (lowercase)
- `%#X` – Base 16 with leading `0X` (uppercase)
- `%4d` &nbsp;– Width 4, right-justified with spaces
- `%-4d` – Width 4, left-justified with spaces
- `%04d` – Width 4, padded with zeroes

---

## String Formatting Verbs

- `%s` &nbsp;– Plain string
- `%q` &nbsp;– Double-quoted string
- `%8s` &nbsp;– Width 8, right-justified
- `%-8s` – Width 8, left-justified
- `%x` &nbsp;– Hex dump of byte values (no spaces)
- `% x` – Hex dump of byte values with spaces

---

## Boolean Formatting Verb

- `%t` – Prints `true` or `false` (same as `%v`)

---

## Float Formatting Verbs

- `%e` &nbsp;– Scientific notation (e.g., `1.23e+03`)
- `%f` &nbsp;– Decimal format, no exponent
- `%.2f` – Precision of 2 digits
- `%6.2f` – Width 6, precision 2 (right-justified)
- `%g` &nbsp;– Uses `%e` or `%f` as needed for compact representation
