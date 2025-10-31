package main

import (
	"fmt"
	"math/rand"
)

func checkAge(age int) {
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}
}

func randomNumber() {
	n := rand.Intn(10)

	if n == 0 {
		fmt.Printf("%d is a low number", n)
	} else if n > 5 {
		fmt.Printf("%d is a big number", n)
	} else {
		fmt.Printf("%d is a good number", n)
	}
}

// Scoping a variable to an if statement
func randomNumberv2() {

	if n := rand.Intn(10); n == 0 {
		fmt.Printf("%d is a low number", n)
	} else if n > 5 {
		fmt.Printf("%d is a big number", n)
	} else {
		fmt.Printf("%d is a good number", n)
	}

}

func CalculateResellPrice(originalPrice float64, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.80
	} else if age >= 10 {
		return originalPrice * 0.50
	} else {
		return originalPrice * 0.70
	}
}

func main() {
	checkAge(16)
	checkAge(18)
	checkAge(8)
	randomNumber()
	randomNumberv2()

	fmt.Println(CalculateResellPrice(1000, 1))
}
