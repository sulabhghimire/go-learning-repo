# Writing Files in Go

## Overview

The `os` package in Go provides utilities to create, open, write, and manage files. Writing to files is a fundamental operation for many applications, and Go offers a clean and efficient API to handle it.

---

## Common Functions

### 📄 File Creation & Opening

- **`os.Create(name string) (*os.File, error)`**  
  Creates a new file. If the file already exists, it will be truncated.

- **`os.OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)`**  
  Opens or creates a file with specified flags and permissions.  
  Common flags:
  - `os.O_CREATE`: Create a new file if none exists
  - `os.O_APPEND`: Append data to the file
  - `os.O_WRONLY`: Write-only
  - `os.O_RDWR`: Read and write

### ✍️ Writing to Files

- **`file.Write(b []byte) (n int, err error)`**  
  Writes a byte slice to the file.

- **`file.WriteString(s string) (n int, err error)`**  
  Writes a string to the file.

---

## Example

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Create or open a file for writing
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Best practice to close when done

	// Write string to file
	_, err = file.WriteString("Hello, Gopher!\n")
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}

	// Write byte slice
	data := []byte("This is a byte slice.\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing bytes:", err)
		return
	}

	fmt.Println("Data written successfully.")
}
```

---

## Best Practices

### ✅ Error Handling

Always check for errors after file operations to avoid silent failures.

### 🧼 Deferred Closing

Use defer file.Close() immediately after opening a file to ensure resources are properly released.

### 🔐 File Permissions

Be cautious with permission flags when using os.OpenFile. Use 0666 for general read-write access (subject to umask).

### 🚀 Buffering (Optional)

For larger writes or performance-critical apps, consider using bufio.Writer for buffered output.
