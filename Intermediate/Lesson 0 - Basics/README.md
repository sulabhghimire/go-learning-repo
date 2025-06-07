# Understanding byte, []byte, rune and string in Go

---

1. In Go byte and rune aren't distinct type but they are alias for integer types.
2. Strings is defiend as immutable sequence of bytes in Go
3. A []byte is a slice of bytes.

---

| Type   | Underlying    | Size      | Mutability     | Description                                                                                     |
| ------ | ------------- | --------- | -------------- | ----------------------------------------------------------------------------------------------- |
| byte   | uint8 (alias) | 1 byte    | Mutable(value) | Represents 8-bit unsigned integer (often for raw data or ASCII Characters)                      |
| rune   | int32 (alias) | 4 bytes   | Mutable(value) | Represents a Unicode code (i.e a character) Commonly used for text                              |
| string | defined type  | 2 words\* | Immutable      | A read-only sequence of bytes. By convention holds UTF-8 encoded text.                          |
| []byte | slice of byte | 3 words\* | Mutable slice  | A mutable, growable sequence of bytes (backed by an array). Often used for I/O and binary data. |

---

## Typical Use Cases

- string (immutable, text): Use for textual data, program messages, keys in maps, or any time you want a safe read-only string. Since strings are immutable and interned, they are efficient for comparison and safe to share. Because strings cannot be modified, operations that “change” a string actually create a new one.
- []byte (mutable, binary): Use for binary data, I/O buffers, and when you need to build or modify byte sequences. A []byte slice can be appended to and modified in place. For example, bytes.Buffer or strings.Builder often accumulate bytes into a slice before converting to string.
- rune (character processing): Use when you need to work with Unicode code points or perform character-wise operations. For example, using the unicode package on rune values (like unicode.IsLetter(r)) ensures correct behavior for multibyte characters.

---

## Conversions Between Types

- String → []byte: b := []byte(s) copies all bytes of string s into a new byte slice.
- []byte → String: s := string(b) copies all bytes of slice b into a new string.
- String → []rune: runes := []rune(s) decodes the UTF-8 string into a slice of runes (code points). Each element is a rune value.
- []rune → String: s := string(runes) builds a string by UTF-8 encoding each rune.
- rune → String (single char): s := string(r) creates a string of length 1 (or more, if r is outside BMP) containing that code point.
