# Go

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