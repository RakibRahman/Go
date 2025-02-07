package main

import "fmt"

// Identifying common behavior among objects and creating a base structure or interface to represent it.

// In Go, generalization is achieved using interfaces  or embedding  (composition).

// Generalization: Define an interface for common behavior
type Vehicle interface {
	Start()
	Stop()
}

// Car implements the Vehicle interface
type Car struct{}

func (c *Car) Start() {
	fmt.Println("Car started")
}

func (c *Car) Stop() {
	fmt.Println("Car stopped")
}

// Bike implements the Vehicle interface
type Bike struct{}

func (b *Bike) Start() {
	fmt.Println("Bike started")
}

func (b *Bike) Stop() {
	fmt.Println("Bike stopped")
}

func main() {
	var vehicle1 Vehicle = &Car{}
	var vehicle2 Vehicle = &Bike{}

	vehicle1.Start()
	vehicle1.Stop()
	vehicle2.Start()
	vehicle2.Stop()
}
