# Hashing

## What is Hashing?

Hashing is a process used in computing to transform data into a fixed-size string of characters, which typically appears random. This transformation is performed using a **hash algorithm**.

---

## Characteristics of Hashing

- 🔒 **Fixed-Size Output**  
  Regardless of input size, the output is of a fixed size.  
  _(e.g., SHA-256 always produces a 256-bit or 32-byte hash)_

- 🔁 **Deterministic**  
  The same input will always produce the same hash output.

- ⚡ **Avalanche Effect**  
  A small change in the input drastically changes the hash output.

- 🔐 **Irreversible**  
  Hash functions are one-way functions; it’s computationally infeasible to reverse-engineer the original input from its hash.

- ⚙️ **Efficient**  
  Hashing operations are designed to be fast and scalable.

---

## Why Hashing?

- ✅ **Secure Data Storage**  
  Commonly used for safely storing passwords in databases.

- 🔍 **Data Integrity Verification**  
  Helps to verify whether data has been tampered with by comparing hash values.

- 🚀 **Fast Data Retrieval**  
  Frequently used in hash tables and maps for quick lookup operations.

---

## Irreversibility of Hashing

In general, **you cannot determine the original input** from its hash value due to the nature of hash functions.

### Why Hashing is Irreversible:

- Fixed-size output for any input size.
- Information is lost during transformation.
- Designed to minimize collisions (different inputs yielding same output).
- Avalanche effect makes reversing unpredictable and infeasible.

> 🔢 Example:  
> SHA-256 produces a 256-bit hash — that’s 2²⁵⁶ possible values!

---

## Salting

**Salting** adds a unique, random value (called a salt) to input before hashing.

### Why Salting is Important:

- Prevents **dictionary attacks**
- Protects against **rainbow table attacks**
- Makes precomputed hash attacks ineffective

Salts can be randomly generated and stored along with the hash for verification.

---

## Hashing in Go (Golang)

Go provides robust hashing utilities in the `crypto` package:

- **SHA-256** → `crypto/sha256`
- **SHA-512** → `crypto/sha512`

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := "hello world"
	hash := sha256.Sum256([]byte(data))
	fmt.Printf("SHA256: %x\n", hash)
}
```
