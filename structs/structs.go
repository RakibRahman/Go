package main

import "fmt"

type person struct {
	name      string
	age       int
	isMarried bool
}

// Anonymous Structs
var person2 struct {
	name string
	age  int
	pet  string
}

var rakib person

func main() {
	zakir := person{"Zakir", 40, true}

	// Anonymous Struct
	pet := struct {
		name string
		kind string
	}{
		name: "Bilai",
		kind: "Cat",
	}

	fmt.Println(rakib)
	rakib.name = "Rakib"
	fmt.Println(zakir)
	fmt.Println(rakib.name)
	person2.name = "Saleh"
	fmt.Println(person2)
	fmt.Println(pet)
	fmt.Println(rakib == zakir)

}
