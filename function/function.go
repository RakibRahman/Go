package main

import (
	"errors"
	"fmt"
	"sort"
)

func add(x int, y int) int {
	return x + y
}

func greet(name string) {
	fmt.Println(("Hello " + name))
}

// Multiple Return Values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}

func divAndReminder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("can not divide by zero")
	}
	return num / denom, num % denom, nil
}

// Named Return Values
func substract(a, b int) (result int) {
	result = a - b
	return
}

// Variadic Functions- accept a variable number of arguments.
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	// for i := 0; i < len(vals)-1; i++ {
	// 	out = append(out, out[i]+base)
	// }
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// named parameters
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func PersonInfo(info Person) string {
	return fmt.Sprintf("Hello, my name is %s %s and I am %d years old.", info.FirstName, info.LastName, info.Age)
}

// function variable
var myFuncVariable func(string) int

func f1(a string) int {
	return len(a)
}

// passing function as parameter
func passingAsParameter() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	// sort by last name
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	// sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a
}

// calling a function doesn’t modify the variable whose value was passed in (unless the variable is a slice or map)
func modifyFails(i int, s string, p Person) {
	i = i * 2
	s = "Goodbye"
	p.FirstName = "Bob"
}

const (
	defaultPreparationTimePerLayer = 2

	noodleQtyPerLayer = 50

	sauceQtyPerLayer = 0.2

	defaultServingInRecipe = 2
)

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = defaultPreparationTimePerLayer
	}
	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	var noodlesNeeded int
	var sauceNeeded float64

	for _, v := range layers {
		if v == "noodles" {
			noodlesNeeded += noodleQtyPerLayer
		}

		if v == "sauce" {
			sauceNeeded += sauceQtyPerLayer
		}
	}

	return noodlesNeeded, sauceNeeded
}
func AddSecretIngredient(secretLayers, layers []string) {
	secretIngredient := secretLayers[len(secretLayers)-1]
	layers[len(layers)-1] = secretIngredient
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	if quantities == nil {
		return nil
	}

	// Create a new slice to avoid modifying the original
	scaled := make([]float64, len(quantities))

	scaleFactor := float64(portions) / defaultServingInRecipe

	for i, v := range quantities {
		scaled[i] = v * scaleFactor
	}
	return scaled
}

func main() {
	//fmt.Println(add(10, 12))
	//greet("Rakib")
	//result1, err := divide(50, 5)
	//result2, err2 := divide(120, 0)
	//
	//fmt.Println(result1, err)
	//fmt.Println(result2, err2)
	//fmt.Println(substract(50, 40))
	//fmt.Println(sum(1, 2, 3, 4, 5)) // Outputs: 15
	//
	//// Anonymous Functions
	//checkAge := func(age int) string {
	//	if age >= 18 {
	//		return "Adult Person"
	//	} else {
	//		return "Minor Person"
	//	}
	//}
	//fmt.Println(checkAge(20))
	//defer fmt.Println("Goodbye!")
	//fmt.Println("Hello")
	//rakib := PersonInfo(Person{
	//	FirstName: "Rakib",
	//	LastName:  "Talukder",
	//	Age:       17,
	//})
	//
	//labib := PersonInfo(Person{
	//	FirstName: "Labib",
	//	Age:       17,
	//})
	//
	//fmt.Println(rakib)
	//fmt.Println(labib)
	//fmt.Println(addTo(3))
	//fmt.Println(addTo(3, 2))
	//fmt.Println(addTo(3, 2, 4, 6, 8))
	//arr := []int{9, 6, 3}
	//fmt.Println(addTo(2, arr...))
	//fmt.Println(divAndReminder(10, 2))
	//result, remainder, err3 := divAndReminder(50, 9)
	//fmt.Println(result, remainder, err3)
	//
	//myFuncVariable = f1
	//res := myFuncVariable("Hello")
	//fmt.Println(res)
	//passingAsParameter()
	//deferExample()
	//p := Person{}
	//i := 2
	//s := "Hello"
	//modifyFails(i, s, p)
	//fmt.Println(i, s, p)
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	fmt.Println(PreparationTime(layers, 3))
	fmt.Println(PreparationTime(layers, 0))
	fmt.Println(Quantities(layers))
	AddSecretIngredient(friendsList, myList)
	fmt.Println(myList)
	quantities := []float64{1.2, 3.6, 10.5}
	scaledQuantities := ScaleRecipe(quantities, 4)
	fmt.Println(scaledQuantities)
}
