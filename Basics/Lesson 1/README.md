# Go Run vs Go Build

This guide explains the difference between `go run` and `go build`, two essential commands for compiling and executing Go programs.

---

## 🚀 Running a Go Program

### `go run filename.go`

- Compiles and executes a Go program in one step.
- **No persistent executable** is created.
- Internally:
  1. The `.go` file is compiled into a temporary executable.
  2. The executable runs immediately.
  3. The temporary executable is discarded after execution.

✅ Useful for:

- Quickly testing or running small programs.
- Development and debugging.

---

## 🔨 Building a Go Program

### `go build filename.go`

- Compiles the Go program into a **standalone executable**.
- Produces a binary file in the current directory (e.g., `filename` or `filename.exe` on Windows).

```bash
go build filename.go
./filename   # Run the compiled executable
```
