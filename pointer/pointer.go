package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Field1 string
	Field2 int
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
}
