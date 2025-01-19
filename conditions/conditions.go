package main

import "fmt"

func checkAge(age int) {
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}
}

func main() {
	checkAge(16)
	checkAge(18)
	checkAge(8)
}
