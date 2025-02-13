package main

import "fmt"

func main() {
	var nilMap map[string]int
	// nilMap["two"] = 2 //will cause panic
	totalWins := map[string]int{}
	totalWins["player_one"] = 111 // can read and write to a map assigned an empty map literal.
	schoolInfo := map[string][]string{
		"teachers": {"Zakir", "Asif"},
		"students": {"labib", "yousha"},
	}
	fmt.Println(nilMap)
	fmt.Println(totalWins)
	fmt.Println(schoolInfo)

}
