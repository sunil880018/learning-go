# Go Project Development Guide

When developing a project in Go, several commands and tools are available to help manage your code, dependencies, and build processes. Here’s a list of essential Go commands commonly used in project development:

## Basic Go Commands

### 1. Create a New Module:

```sh
go mod init <module-name>
```

Initializes a new Go module, creating a `go.mod` file to manage dependencies.

### 2. Build the Project:

```sh
go build
```

Compiles the code in the current directory (or the specified package) and produces an executable binary.

### 3. Run the Project:

```sh
go run <file.go>
```

Compiles and runs the specified Go file. You can also specify multiple files.

### 4. Test the Project:

```sh
go test
```

Runs tests in the current package. You can also specify a package or file:

```sh
go test <package>
```

or

```sh
go test <file_test.go>
```

### 5. Install a Package:

```sh
go get <package>
```

Downloads and installs the specified package and its dependencies.

### 6. Install Dependencies:

```sh
go mod tidy
```

This command does the following:

- Adds any missing modules necessary to build the current module and its dependencies.
- Removes any modules that are no longer necessary.
- Updates the `go.sum` file with checksums for new dependencies.

### 7. Update Dependencies:

```sh
go get -u
```

Updates the dependencies in the module to their latest versions.

### 8. List Dependencies:

```sh
go list ...
```

Lists all the packages in the module, including dependencies.

### 9. Format the Code:

```sh
go fmt
```

Formats Go source code files in the current directory according to the Go style guidelines.

### 10. Generate Documentation:

```sh
go doc <package>
```

Displays documentation for the specified package or function.

### 11. Check for Errors:

```sh
go vet
```

Analyzes your code for potential errors and issues that are not caught by the compiler.

### 12. Run a Specific Test:

```sh
go test -run TestFunctionName
```

Runs a specific test function.

### 13. Benchmark Tests:

```sh
go test -bench=.
```

Runs all benchmark tests in the current package.

### 14. Clean Up Build Artifacts:

```sh
go clean
```

Removes object files and cached files.

---

This guide provides the fundamental Go commands needed for project development. Happy coding! 🚀

## Managing Dependencies (Equivalent to npm commands):

| Node.js (npm/yarn)          | Go (go mod)             |
| --------------------------- | ----------------------- |
| `npm init`                  | `go mod init <module>`  |
| `npm install`               | `go get <package>`      |
| `npm update`                | `go get -u`             |
| `npm install package@1.0.0` | `go get package@v1.0.0` |
| `npm list`                  | `go list -m all`        |
| `package-lock.json`         | `go.sum`                |
| `nodemon`                   | `air`                   |

Would you like an example project structure in Go similar to a Node.js project?

### like nodemon in node js, similar in go (Air)

### 1.install air

```sh
 go install github.com/air-verse/air@latest

```

### 2.set path

```sh
open nano ~/.zshrc
```

```sh
export PATH=$PATH:$(go env GOPATH)/bin
```

### 3.Save and exit

```sh
In nano: press CTRL+O → Enter → CTRL+X
```

### 4. Reload .zshrc

```sh
source ~/.zshrc
```

### 5. Now check if air works

```sh
air -v
```

### 6. Run the main.go file

```sh
air
```
