package main

import "fmt"

type Vehicle interface {
	Start()
	Stop()
}

// hidden implemention
type Engine struct{}

func (e *Engine) Ignite() {
	fmt.Println("Engine Ignited")
}

func (e *Engine) Shutdown() {
	fmt.Println("Engine shutdown")
}

// Car struct (exposes only Start and Stop)
type Car struct {
	engine *Engine
}

func (c *Car) Start() {
	c.engine.Ignite()
	fmt.Println("Car started")
}

func (c *Car) Stop() {
	c.engine.Shutdown()
	fmt.Println("Car stoped")
}

func main() {
	engine := &Engine{}
	car := Car{engine: engine}

	car.Start()
	car.Stop()
}
