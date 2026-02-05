package main

import (
	"fmt"
	"math"
)

type Stringer interface {
	String() string
}

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// business logic
	return ""
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	// get data from somewhere
	c.L.Process("data")
}

type Speaker interface {
	Speak() string
}

type Cat struct{}

func (d Cat) Speak() string { return "Meow!" } // Cat now implements Speaker implicitly.

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64 { return r.Width * r.Height }

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

// Type Assertions
func typeAssertionDemo() {
	var i interface{} = "World"
	str, ok := i.(string)
	if ok {
		fmt.Println("String value:", str)
	} else {
		fmt.Println("Not a string")
	}
	var data interface{} = map[string]interface{}{
		"name": "Nusaiba",
		"age":  30,
	}
	if m, ok2 := data.(map[string]interface{}); ok2 {
		fmt.Println("Name:", m["name"])
	}
}

// Type Conversion
func typeConvertionDemo() {
	var x int32 = 42
	var y int64 = int64(x) // Convert int32 to int64
	fmt.Println(y)         // Output: 42
}

// Type Switch
func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("It's an integer:", v)
	case string:
		fmt.Println("It's a string:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
	circleShape := Circle{radius: 90}
	rectShape := Rectangle{Width: 20, Height: 20}
	PrintArea(circleShape)
	PrintArea(rectShape)
	typeAssertionDemo()
	describe(42)
	typeConvertionDemo()
}
