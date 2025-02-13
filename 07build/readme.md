# Go Build Commands Explained

This document explains different ways to use `go build` to compile Go programs.

## 1. `go build main.go`

- Compiles only the `main.go` file.
- Does **not** produce an output file in the current directory.
- Instead, it compiles and links the binary but deletes it immediately after checking for errors.
- Useful for quick compilation checks.

## 2. `go build`

- Compiles all Go files in the current package (directory).
- If the package is a `main` package, it produces an executable named after the directory.
- If it's not a `main` package, it does nothing (since libraries don’t generate standalone executables).

## 3. `go build -o hello-world`

- Compiles the current package and outputs an executable named `hello-world`.
- Useful when you want to specify a custom name for the binary.

## 4. `go build -o hello-world main.go`

- Explicitly compiles only `main.go` and outputs an executable named `hello-world`.
- Similar to `go build -o hello-world`, but ensures only `main.go` is compiled.

---

## 5. Cross-Compilation with `GOOS`

Go allows cross-compiling binaries for different operating systems using the `GOOS` environment variable.

### Compile for Windows

```sh
GOOS=windows go build -o myapp.exe
```

- Generates a Windows-compatible `.exe` binary.
- Useful when building on a non-Windows machine but targeting Windows.

### Compile for Linux

```sh
GOOS=linux go build -o myapp
```

- Generates a Linux-compatible binary.
- Can be run on Linux servers or cloud platforms.

### Compile for macOS (Darwin)

```sh
GOOS=darwin go build -o myapp
```

- Generates a macOS-compatible binary.
- Useful for Apple devices running macOS.

---

## Summary Table

| Command                              | Compiles? | Produces Executable?   | Output Name    |
| ------------------------------------ | --------- | ---------------------- | -------------- |
| `go build main.go`                   | ✅        | ❌ (temp binary)       | N/A            |
| `go build`                           | ✅        | ✅ (if `main` package) | Directory name |
| `go build -o hello-world`            | ✅        | ✅                     | `hello-world`  |
| `go build -o hello-world main.go`    | ✅        | ✅                     | `hello-world`  |
| `GOOS=windows go build -o myapp.exe` | ✅        | ✅                     | `myapp.exe`    |
| `GOOS=linux go build -o myapp`       | ✅        | ✅                     | `myapp`        |
| `GOOS=darwin go build -o myapp`      | ✅        | ✅                     | `myapp`        |

---

## Usage Examples

### Build and Run in Current OS

```sh
go build
./your-executable-name
```

### Build with Custom Output Name

```sh
go build -o myapp
./myapp
```

### Cross-Compile for Different Operating Systems

```sh
GOOS=windows go build -o myapp.exe
GOOS=linux go build -o myapp
GOOS=darwin go build -o myapp
```

Now you can distribute your Go binaries across different platforms! 🚀
