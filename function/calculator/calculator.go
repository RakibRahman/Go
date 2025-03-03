package main

import (
	"errors"
	"fmt"
)

type OperationFunc func(int, int) (int, error)

func add(i, j int) (int, error) { return i + j, nil }

func sub(i, j int) (int, error) { return i - j, nil }

func mul(i, j int) (int, error) { return i * j, nil }

func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

var operationMap = map[string]OperationFunc{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	op := operationMap["+"]
	subOp := operationMap["-"]
	multiOp := operationMap["*"]
	divOp := operationMap["/"]

	fmt.Println(op(1, 2))
	fmt.Println(subOp(96, 9))
	fmt.Println(multiOp(96, 9))
	fmt.Println(divOp(90, 3))
}
