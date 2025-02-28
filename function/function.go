package main

import (
	"errors"
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func greet(name string) {
	fmt.Println(("Hello " + name))
}

// Multiple Return Values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}

func divAndReminder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("can not divide by zero")
	}
	return num / denom, num % denom, nil
}

// Named Return Values
func substract(a, b int) (result int) {
	result = a - b
	return
}

// Variadic Functions- accept a variable number of arguments.
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	// for i := 0; i < len(vals)-1; i++ {
	// 	out = append(out, out[i]+base)
	// }
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// named parameters
type Info struct {
	FirstName string
	LastName  string
	Age       int
}

func PersonInfo(info Info) string {
	return fmt.Sprintf("Hello, my name is %s %s and I am %d years old.", info.FirstName, info.LastName, info.Age)
}

// function variable
var myFuncVariable func(string) int

func f1(a string) int {
	return len(a)
}

func main() {
	fmt.Println(add(10, 12))
	greet("Rakib")
	result1, err := divide(50, 5)
	result2, err2 := divide(120, 0)

	fmt.Println(result1, err)
	fmt.Println(result2, err2)
	fmt.Println(substract(50, 40))
	fmt.Println(sum(1, 2, 3, 4, 5)) // Outputs: 15

	// Anonymous Functions
	checkAge := func(age int) string {
		if age >= 18 {
			return "Adult Person"
		} else {
			return "Minor Person"
		}
	}
	fmt.Println(checkAge(20))
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello")
	rakib := PersonInfo(Info{
		FirstName: "Rakib",
		LastName:  "Talukder",
		Age:       17,
	})

	labib := PersonInfo(Info{
		FirstName: "Labib",
		Age:       17,
	})

	fmt.Println(rakib)
	fmt.Println(labib)
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	arr := []int{9, 6, 3}
	fmt.Println(addTo(2, arr...))
	fmt.Println(divAndReminder(10, 2))
	result, remainder, err3 := divAndReminder(50, 9)
	fmt.Println(result, remainder, err3)

	myFuncVariable = f1
	res := myFuncVariable("Hello")
	fmt.Println(res)
}
