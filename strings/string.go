package main

import (
	"fmt"
	"strings"
)

func CleanupMessage(oldMsg string) string {
	removeStarsFromStr := strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(removeStarsFromStr)
}

func main() {
	var name = "rakib"
	age := 28
	names := "John doe matthew"
	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`
	var firstChar byte = name[0] //returns 114 (the UTF-8 value of 'r')
	fmt.Println(firstChar)
	var firstChar2 string = name[0:1] //returns 'r'
	fmt.Println(firstChar2)
	var s string = "Hello 👍"
	fmt.Println(len(s))
	formattedString := fmt.Sprintf("My name is %s and i'm %d years old", name, age)
	fmt.Println(formattedString)
	fmt.Println(len(name))
	fmt.Println(strings.Contains(name, "r"))
	fmt.Println(strings.Split(names, ","))
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
	fmt.Println(CleanupMessage(message))
}
