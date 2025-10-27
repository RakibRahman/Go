package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {

	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 2, Height: 2}

	// Go uses dynamic dispatch here — it calls the correct method based on the actual type at runtime.
	printArea(c)
	printArea(r)
}
