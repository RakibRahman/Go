package main

import "fmt"

// A strong relationship where one object is composed of another, and the contained object cannot exist without the container.

type Room struct {
	Name string
}

type House struct {
	Rooms []Room
}

// constructor function
func NewHouse() *House {
	return &House{
		Rooms: []Room{
			{Name: "Living Room"},
			{Name: "Bedroom"},
			{Name: "Kitchen"},
		},
	}
}

func main() {
	house := NewHouse()
	for _, room := range house.Rooms {
		fmt.Println("-", room.Name)
	}
	// If the house is destroyed, the rooms are also destroyed
	house = nil
}
