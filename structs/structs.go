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

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

func NewTrack(distance int) Track {
	return Track{distance}
}

func Drive(car Car) Car {

	if car.battery < car.batteryDrain {
		return car
	}

	newBattery := car.battery - car.batteryDrain
	newDistance := car.distance + car.speed

	return Car{
		battery:      newBattery,
		batteryDrain: car.batteryDrain,
		speed:        car.speed,
		distance:     newDistance,
	}
}

func CanFinish(car Car, track Track) bool {
	maxDrives := car.battery / car.batteryDrain
	maxDistances := maxDrives * car.speed

	return maxDistances >= track.distance
}

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
	speed := 5
	batteryDrain := 2
	distance := 100
	track := NewTrack(distance)
	car := NewCar(speed, batteryDrain)
	car = Drive(car)
	fmt.Println(car)
	fmt.Println(track)
	fmt.Println(CanFinish(car, track))

}
