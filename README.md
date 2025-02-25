# Go
Go is a small language with only 25 keywords.

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

```
 ### String  

- A string in Go is a sequence of bytes that represents text.  
- Strings are immutable, meaning their content cannot be changed once created.  
- Internally, Go stores strings as UTF-8 encoded byte sequences.  

### Byte
- A byte is an 8-bit unsigned integer (uint8).  

### Rune
- A rune is an alias for int32 and represents a Unicode code point.  

# Struct
Maps are a convenient way to store some kinds of data, but they have limitations.
They don’t define an API since there’s no way to constrain a map to allow only certain
keys. Also, all values in a map must be of the same type. For these reasons, maps are
not an ideal way to pass data from function to function.


# Blocks
Each place where a declaration occurs is called a block. Variables, constants, types,
and functions declared outside of any functions are placed in the package block.
All the variables defined at the top level of a function
(including the parameters to a function) are in a block. Within a function, every set of
braces ({}) defines another block.

# The Universe Block
The built-in types (like int and string), constants (like true and false), and functions  (like make or close) are considered as 
predeclared identifiers and defines them in the universe block, which is the block that contains all other blocks.

# Shadowing Variables
A shadowing variable is a variable that has the same name as a variable in a containing
block. For as long as the shadowing variable exists, you cannot access a shadowed
variable.
```
func main() {
x := 10
fmt.Println(x)
fmt := "oops" //shadows fmt package in the file block
fmt.Println(fmt) // give error, as local variable `fmt` does not have reqwuired methods
}
```