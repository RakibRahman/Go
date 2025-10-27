package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cant divide by zero")
	}
	return a / b, nil
}

// Custom errors
type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}

func doSomething() error {
	return MyError{Code: 404, Msg: "Not Found"}
}

var ErrorNotFound = errors.New("not found")

func findItem() error {
	return fmt.Errorf("operation failed: %w", ErrorNotFound)
}

func main() {

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Result:", result)

	err2 := doSomething()

	if err2 != nil {
		fmt.Println(err) // Error 404: Not Found
	}

	err3 := findItem()
	if errors.Is(err3, ErrorNotFound) {
		fmt.Println("Item not found")
	}
}
