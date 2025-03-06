# Go
Go is a small language with only 25 keywords.

# Call by value
In Go, function arguments are passed by value, meaning that a copy of the argument's value is made and passed to the function. This means that any changes made to the parameter within the function do not affect the original argument.

In Go, maps and slices are reference types, meaning that when you pass a map or a slice to a function, you're passing a reference to the underlying data, not a copy of the data. This allows the function to modify the original data directly.

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

# Closures
Functions declared inside functions are special; they are closures. This is a computer
science word that means that functions declared inside functions are able to access
and modify variables declared in the outer function.

`Go` also uses **closures** to implement resource cleanup, via the `defer` keyword.

You can defer multiple functions in a Go function. They run in last-in, first-out
(LIFO) order; the last defer registered runs first.

```
func main() {
a := 20
f := func() {
fmt.Println(a)
a := 30
fmt.Println(a)
}
f()
fmt.Println(a)
}

output:
20
30
20
```

# Functions Versus Methods
## Function
- A function in go is a standalone code block which performs a specific task.
- Functions  are general and unattached to types.
```
  Pseudo Code:
  func FunctionName(parameters) returnType {
    // Code logic
}

func sum(x,y int) int{
    return x,y
}
```

 ## Methods:
  - Function works independently and does not relies on any particular data type like struct or custom type.

  - A method is a function with a receiver which is associated with specifc data type, like struct or custom type.
   receiver allows a method to perform tasks within defined tyoe behavior.

```
   Pseudo Code:
   func (receiverType Receiver) MethodName(parameters) returnType {
    // Code logic
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	// receiver - (p Person)
	// Method name - String()
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
```
Methods  bind behavior to a type via a receiver, enabling encapsulation and type-specific logic.
Methods enable encapsulation and "object-oriented" patterns in Go, while functions handle general logic. Use methods to bind behavior to data, and functions for reusable operations. 
     

# Interfaces
In Go, interfaces  define a set of method signatures, acting as a contract for behavior.
A type implicitly  implements an interface by implementing all its methods—no explicit declaration required. This promotes clean, decoupled code.  
- Interfaces specify what callers need. The client code defines the interface to specify what functionality it requires.
- Interfaces are checked at compile time, which helps catch errors early.
- In the Go runtime, interfaces are implemented as a struct with two pointer fields, one for the value and one for the type
of the value. an interface is equal to nil only if its type and value fields are both nil
- An empty interface typesimply states that the variable can store any value whose type implements zero or more methods.
- Accept Interfaces, Return Structs


# Type Assertions
Type assertions are used to extract the underlying concrete value of an interface type. They allow you to "assert" that an interface holds a specific type.  
# Type Conversion
Type conversion explicitly converts a value from one type to another. This is only allowed between compatible types (e.g., numeric types, custom types based on the same underlying type).
Type conversion does not  work between unrelated types (e.g., int to string).


A type assertion is very different from a type conversion. Conversions change a value to a new type, while assertions reveal the
type of the value stored in the interface. Type conversions can be applied to both concrete types and interfaces. 
Type assertions can be applied only to interface types. All type assertions are checked at runtime, so they can fail at runtime with a panic if you don’t
use the comma ok idiom.

# Type Switch
handle multiple possible types
```
func doThings(i any) {
	switch j := i.(type) {
	case nil:
	// i is nil, type of j is any
	case int:
	// j is of type int
	case MyInt:
	// j is of type MyInt
	case io.Reader:
	// j is of type io.Reader
	case string:
	// j is a string
	case bool, rune:
	// i is either a bool or rune, so j is of type any
	default:
		// no idea what i is, so j is of type any
	}
}
```