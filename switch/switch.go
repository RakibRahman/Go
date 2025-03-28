package main

import (
	"fmt"
	"math/rand"
)

func randomNumbers() {
	numbers := make([]int, 0, 100)
	for i := 0; i <= 100; i++ {
		num := rand.Intn(100)
		numbers = append(numbers, num)
		switch {
		case num%6 == 0:
			fmt.Println("SIx")
		case num%2 == 0:
			fmt.Println("Two")
		case num%3 == 0:
			fmt.Println("Three")
		default:
			fmt.Println("Ignore")
		}
	}
}

func main() {
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word")
		case 5:
			fmt.Println(word, "is exactly the right length:", size)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	//blank switch
	a := 2
	switch {
	case a == 2:
		fmt.Println("a is 2")
	case a == 3:
		fmt.Println("a is 3")
	case a == 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("a is ", a)
	}
	a = 4
	switch a {
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	case 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("a is ", a)
	}

	randomNumbers()

}
