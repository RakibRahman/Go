package main

import "fmt"

func main() {
	var name string
	var age byte = 0
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)
	fmt.Printf("Hello %s, Your age is %d", name, age)
}
