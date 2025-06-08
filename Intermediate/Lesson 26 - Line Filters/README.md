# Line Filtering

## What is Line Filtering?

**Line Filtering** refers to the process of processing or modifying lines of text based on specific criteria. It typically involves reading input line by line and applying a transformation or condition to determine whether a line should be retained, altered, or discarded.

This technique is commonly used in:

- ✅ Text processing
- 🧹 Data cleaning
- 📂 File manipulation

---

## Examples of Line Filtering

Some common use cases include:

- 🔍 **Filtering lines based on content**  
  e.g., only keep lines containing the word `"error"`

- 🧼 **Removing empty lines**

- ✏️ **Transforming line contents**  
  e.g., trimming whitespace, converting to lowercase

- 📏 **Filtering by length**  
  e.g., printing lines longer than a certain number of characters

---

## Best Practices

1. **Use Buffered Input**  
   Use buffered readers (e.g., `bufio.Scanner` in Go) to efficiently handle large files line by line.

2. **Error Handling**  
   Always check and handle I/O errors while reading or processing lines.

3. **Input Source Flexibility**  
   Design your solution to accept input from various sources (e.g., stdin, file, network).

4. **Flexibility**  
   Write modular code so filtering conditions or transformations can be easily extended or configured.

---

## Practical Applications

- 📊 **Data Transformations**  
  Cleaning or restructuring input data before importing into a database.

- 📜 **Text Processing**  
  Parsing logs, filtering documentation, formatting text for display.

- 📈 **Data Analysis**  
  Selecting relevant rows from large datasets for further processing or visualization.

---

## Example in Go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Example: Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Example: Filter lines containing "Go"
		if strings.Contains(line, "Go") {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scan error:", err)
	}
}
```
