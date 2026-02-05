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

func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10

	default:
		return 0
	}
}

func FirstTurn(card1, card2, dealerCard string) string {
	player1 := ParseCard(card1)
	player2 := ParseCard(card2)
	dealer := ParseCard(dealerCard)
	playerScore := player1 + player2

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case playerScore == 21:
		if dealer == 10 || dealer == 11 {
			return "S"
		}
		return "W"
	case playerScore >= 17:
		return "S"
	case playerScore <= 11:
		return "H"
	default:
		if dealer >= 7 {
			return "H"
		}
		return "S"

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
	fmt.Println(ParseCard("ace"))   // 11
	fmt.Println(ParseCard("seven")) // 7
	fmt.Println(ParseCard("king"))  // 10
	fmt.Println(ParseCard("joker")) // 0
	fmt.Println(FirstTurn("ace", "ace", "jack"))
	fmt.Println(FirstTurn("ace", "king", "ace"))
	fmt.Println(FirstTurn("five", "queen", "ace"))
}
