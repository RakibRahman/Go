package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = "rakib"
	age := 28
	names := "John doe matthew"

	formattedString := fmt.Sprintf("My name is %s and i'm %d years old", name, age)
	fmt.Println(formattedString)
	fmt.Println(len(name))
	fmt.Println(strings.Contains(name, "r"))
	fmt.Println(strings.Split(names, ","))

}
