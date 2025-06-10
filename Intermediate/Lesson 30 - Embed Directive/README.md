# 📦 Go `embed` Directive

The `embed` directive was introduced in **Go version 1.16** to embed static files or directories into Go binaries at build time.  
It provides a convenient and effective way to include assets directly within Go programs, eliminating the need to manage these assets separately.

---

## ✅ Why Use the `embed` Directive?

1. **Simplicity**

   - Embedding files simplifies deployment by reducing the number of separate files to manage.

2. **Efficiency**

   - Embedding files into binaries makes distribution and execution straightforward, without worrying about file paths or external dependencies.

3. **Security**
   - Embedded files are bundled within the binary, minimizing exposure to external manipulation or unauthorized access.

---

## 📂 Supported Types

- Individual **files**
- Entire **directories** (must contain at least one embeddable file)

---

## 🚀 Common Use Cases

- Web servers (e.g., serving HTML, CSS, JS)
- Configuration files (e.g., JSON, YAML)
- Testing resources (e.g., test data files)

---

## ⚠️ Considerations

1. **File Size**  
   Large files can significantly increase binary size.

2. **Update Strategy**  
   Embedded content requires recompilation to reflect changes.

3. **Compatibility**  
   Requires **Go 1.16** or later. Older versions do not support the `embed` directive.

---

## 📘 Example

```go
package main

import (
    "embed"
    "fmt"
)

//go:embed config.json
var configFile []byte

func main() {
    fmt.Println(string(configFile))
}
```
