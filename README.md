# Go

# Commands
### GO Module
- Command : `go mod init hello_world`
- A Go project is called a module. A module is not just source code. It is also an exact specification of the dependencies of the code within the module. Every
module has a `go.mod` file in its root directory. Running go mod init creates this file.
- The `go.mod` file declares the name of the module, the minimum supported version of Go for the module, and any other modules that your module depends on

### GO Build
- Command : `go build` / `go build -o hello`
- Creates an executable file (hello) in the current directory.

### GO FMT
- Command: `go fmt ./...` 
- Automatically fixes the whitespace in your code to match the standard format.
-  `./...` tells a Go tool to apply the command to all the files in the current directory and all subdirectories.
    
### GO VET
- Command :  ` go vet ./...` 
- To catches several common programming errors.
- Scan for possible bugs in valid code.


# Variables
Variables are used to store values. It can be declared using the var keyword, or the shorthand := syntax.
```
var name string
name = "John Doe"

age := 30

```


# Variadic Functions
Variadic functions accept a variable number of arguments. Use an ellipsis (...) before the type to indicate that a function takes a variable number of parameters:
```
func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3, 4, 5)) // Outputs: 15
}

```


# Defer Statement
The `defer` statement delays the execution of a function until the surrounding function returns. It’s often used for clean-up tasks:
```
func main() {
    defer fmt.Println("Goodbye!")
    fmt.Println("Hello")
}

``