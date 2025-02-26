package main

import (
	"fmt"
	"math/rand"
)

// The Complete for Statement
func completeFor() {
	for i := 0; i < 10; i++ { // must use := to initialize the variables; var is not legal here
		fmt.Println(i)
	}

	x := 0

	for ; x < 10; x++ {
		fmt.Println(x * 2)
	}

	for y := 0; y < 10; {
		fmt.Println(y)
		if y%2 == 0 {
			y++
		} else {
			y += 2
		}
	}
}

// The Condition-Only for Statement
func conditionOnlyFor() {
	//When you leave off both the initialization and the increment in a for statement, its act like a while
	i := 1
	for i < 100 {
		fmt.Println(i)
		i++
	}
}

// The Infinite for Statement
func infiniteFor() {
	for {
		fmt.Println("Hello")
	}
}

// The for-range Statement
func rangeFor() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		// The first variable is the position in the data structure being iterated, while the second is the value at that position.
		fmt.Println(i, v)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for key := range uniqueNames {
		fmt.Println(key)
	}
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}

// Labeling for Statements
func labelingFor() {
	samples := []string{"hello", "apple_n!"}

outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}

// Break and continue
func breakContinue() {
	evenVals := []int{2, 4, 6, 8, 10}

	//break: exit from a loop prematurely
	for _, v := range evenVals {
		if v == 8 {
			break // Exit the loop when v is 8
		}
		fmt.Println(v)
	}
	fmt.Println("Loop exited.")

	// continue: skip the remaining code inside the loop for the current iteration and move to the next iteration of the loop.
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println(i)
	}
}

func randomNumbers() {
	numbers := make([]int, 0, 100)
	for i := 0; i <= 100; i++ {
		numbers = append(numbers, rand.Intn(100))
	}
	fmt.Println(numbers)
}

func main() {
	// rangeFor()
	// labelingFor()
	// breakContinue()
	randomNumbers()
}
