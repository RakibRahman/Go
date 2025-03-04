package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Field1 string
	Field2 int
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func update(px *int) {
	*px = 20
}

// using pointers

// Don’t do this
func MakeFoo(f *Foo) error {
	f.Field1 = "Hello"
	f.Field2 = 9
	return nil
}

// Do this
func MakeFooV2() (Foo, error) {
	f := Foo{
		Field1: "Rakib",
		Field2: 20,
	}
	return f, nil
}

// Working with JSON
func MakeJSON() {
	payload := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}
	err := json.Unmarshal([]byte(`{"name:"rakib",age:20}`), &payload)
	fmt.Println(err)
}

func UpdateSlice(s []string, val string) {
	index := len(s) - 1
	s[index] = val
	fmt.Println("in UpdateSlice:", s)
}

func GrowSlice(s []string, val string) {
	s = append(s, val)
	fmt.Println("in GrowSlice:", s)
}

func main() {
	var x int32 = 10
	var y bool = true
	z := 101
	pointerX := &x
	pointerY := &y
	fmt.Println(x, y, pointerX, pointerY, *pointerX)

	var f *int //f is nil
	failedUpdate(f)
	fmt.Println(f)
	failedUpdate(&z)
	fmt.Println(z)
	update(&z)
	fmt.Println(z)
	MakeJSON()
	person1 := MakePerson("rakib", "talukder", 28)
	fmt.Println(person1)
	person2 := MakePersonPointer("Wilma", "Fredson", 32)
	fmt.Println(person2)
	s := []string{"a", "b", "c"}
	UpdateSlice(s, "d")
	fmt.Println("in main after UpdateSlice:", s)
	GrowSlice(s, "e")
	fmt.Println("in main, after GrowSlice:", s)
}
