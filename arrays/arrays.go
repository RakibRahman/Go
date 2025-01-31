package main

import "fmt"

func main() {
	var x [3]int
	var y = [3]int{12, 20, 30}
	fmt.Println(x)
	fmt.Println(y)

	//use == and != to compare two arrays. Arrays are equal if they are the same length and contain equal values
	fmt.Println(x == y)
	fmt.Println(x != y)

}
