package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Score int
type Converter func(string) Score
type TeamScores map[string]Score
type Counter struct {
	total       int
	lastUpdated time.Time
}
type Adder struct {
	start int
}

// Methods
func (p Person) String() string {
	// receiver - (p Person)
	// Method name - String()
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// pointer and value receivers

func (c *Counter) Increment() { // use pointer receiver if method is gonna modify the receiver or you methods needs to handle nil instance.
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string { // use value receiver if the method does not modify the receiver.
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	fmt.Println(p) //The String() method is automatically called to format the output.
	output := p.String()
	fmt.Println(output)
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
	c2 := &Counter{total: 10, lastUpdated: time.Now()}
	doUpdateWrong(*c2)
	doUpdateRight(c2)
	myAdder := Adder{start: 85}
	fmt.Println(myAdder.AddTo(96))
	f1 := myAdder.AddTo
	fmt.Println(f1(10))
	f2 := Adder.AddTo // create a function from type itself, its called method expression
	fmt.Println(f2(myAdder, 925))
}
