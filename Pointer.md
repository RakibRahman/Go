# Pointer 

A pointer is a variable that stores the memory address of another variable. A pointer is a variable that holds the location in memory where a value is stored.
Pointers Indicate Mutable Parameters.

if a pointer is passed to a function, the function gets a copy of the pointer.
This still points to the original data, which means that the original data can be
modified by the called function.

Go uses parameters of pointer type to indicate that a
parameter might be modified by the function

If a struct is large enough, using a pointer to the struct as either an input parameter
or a return value improves performance.

In Go, when you pass a map to a function, any modifications made to the map inside the function are reflected in the original map. This behavior occurs because, under the hood, a map in Go is implemented as a pointer to a struct. When you pass a map to a function, you are passing a copy of this pointer, not the actual data structure itself. Since both the original map and the copy of the pointer point to the same underlying data structure, any changes made through one pointer will be visible through the other.
 
 Pointers are useful for: 
- Passing large data structures efficiently (by reference instead of by value).
- Modifying the original variable inside a function.
- Managing memory more explicitly.

- Rather than populating a struct by passing a pointer to it into a function, have the function instantiate and return the struct.

```
// Don’t do this
func MakeFoo(f *Foo) error {
f.Field1 = "val"
f.Field2 = 20
return nil
}

// Do this
func MakeFoo() (Foo, error) {
f := Foo{
Field1: "val",
Field2: 20,
}
return f, nil
}
```

- use pointer parameters to modify a variable is when the
function expects an interface. this pattern is common when working with JSON

# The & Operator (Address-of Operator)
The & operator is used to get the memory address of a variable. It returns a pointer to the variable.
For structs, use an `&` before a struct literal to create a pointer instance.

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


# The * Operator (Indirection/Dereference Operator)
The * operator is used to access the value stored at a memory address (i.e., dereferencing a pointer). It allows you to read or modify the value that the pointer points to.

Before dereferencing a pointer, we must make sure that the pointer is non-nil,otherwise  it will cause panic.

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


```
func UpdateSlice(s []string, val string) {
	index := len(s) - 1
	s[index] = val
	fmt.Println("in UpdateSlice:", s)
}

func GrowSlice(s []string, val string) {
	s = append(s, val)
	fmt.Println("in GrowSlice:", s)
}

output: 
in UpdateSlice: [a b d] // Both slices share the same underlying array.
in main after UpdateSlice: [a b d]

in GrowSlice: [a b d e]
in main, after GrowSlice: [a b d] //A new underlying array is allocated, so the original slice is unchanged.
```


### Slices Share the Same Underlying Array

- A slice is a **descriptor** that contains:
  - A pointer to an underlying array.
  - The length (`len`) of the slice.
  - The capacity (`cap`) of the slice.

- When you pass a slice to a function, a **copy of the slice descriptor** is made, but the **underlying array** is shared between the original slice and the copied slice.

---

### Modifying Elements of a Slice

- If you modify the elements of the slice (e.g., `s[len(s)-1] = val`), the change is reflected in the original slice because both the original slice and the copied slice share the same underlying array.

---

### Appending to a Slice

- Appending to a slice may or may not affect the original slice, depending on whether the capacity of the slice is exceeded:
  - **If the capacity is sufficient**, the new element is added to the existing underlying array, and the change is visible in the original slice.
  - **If the capacity is exceeded**, a new underlying array is allocated, and the original slice remains unchanged because it still points to the old underlying array.
     