package main

import "fmt"

// Controlling access to properties and behaviors using access modifiers (private, protected, public). in Go, with unexxported field

//Struct Person is exported
type Person struct {
	Name string //exported field
	age  int    // Unexported field
}

//Struct company is non-exported
type company struct {
}

//exported method
func (p *Person) GetAge() int {
	return p.age
}

//non-exported
func (p *Person) getName() string {
	return p.Name
}

type car struct {
	engine engine // Unexported field (encapsulated)
}

type engine struct {
	isRunning bool
}

func (e *engine) ignite() {
	e.isRunning = true
	fmt.Println("Engine ignited")
}

func (e *engine) shutdown() {
	e.isRunning = false
	fmt.Println("Engine shutdown")
}

func (c *car) Start() {
	c.engine.ignite()
	fmt.Println("Car started")
}

func (c *car) Stop() {
	c.engine.shutdown()
	fmt.Println("Car stopped")
}

func main() {
	p := &Person{Name: "rakib", age: 23}
	fmt.Println(p)
	c := &company{}
	fmt.Println(c)

	//STRUCTURE'S FIELDS
	fmt.Println(p.Name)
	fmt.Println(p.age)

	//STRUCTURE'S METHOD
	fmt.Println(p.GetAge())
	fmt.Println(p.getName())

	myCar := car{engine: engine{}}
	myCar.Start()
	myCar.Stop()

	// myCar.engine is not accessible here because it's unexported
	fmt.Println(myCar.engine)
}
