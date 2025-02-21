# Pointer 

A pointer is a variable that stores the memory address of another variable. Pointers are useful for: 
- Passing large data structures efficiently (by reference instead of by value).
- Modifying the original variable inside a function.
- Managing memory more explicitly.

# The & Operator (Address-of Operator)
The & operator is used to get the memory address of a variable. It returns a pointer to the variable.

```
pointer := &variable
// &variable returns the memory address where variable is stored

func main() {
    x := 42         // Declare an integer variable `x`
    p := &x         // Get the memory address of `x` using `&`
    
    fmt.Println("Value of x:", x)   // Prints: Value of x: 42
    fmt.Println("Address of x:", p) // Prints the memory address of `x`
}
```


# The * Operator (Dereference Operator)
The * operator is used to access the value stored at a memory address (i.e., dereferencing a pointer). It allows you to read or modify the value that the pointer points to.

```
value := *pointer
// *pointer retrieves the value stored at the memory address held by pointer.


func main() {
    x := 42         // Declare an integer variable `x`
    p := &x         // Get the memory address of `x`
    
    fmt.Println("Value of x:", x)        // Prints: Value of x: 42
    fmt.Println("Address of x:", p)      // Prints the memory address of `x`
    fmt.Println("Value at address:", *p) // Dereference `p` to get the value of `x`
    
    *p = 100        // Modify the value of `x` through the pointer `p`
    fmt.Println("New value of x:", x)    // Prints: New value of x: 100
}

```

# Keypoints:
- Use &x to get the memory address  of x.
- Use *p to get the actual value  stored at the memory address held by p.
- You cannot get the actual value of x directly using &. You must first obtain the pointer (&x) and then dereference it (*p) to access the value.
     