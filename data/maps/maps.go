package main

import (
	"fmt"
	"maps"
)

func mapAsSet() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(intSet)
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}

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
	fmt.Println(len(schoolInfo))
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	m := map[string]int{
		"hello": 5,
		"world": 0,
		"code":  9,
	}

	//The comma ok Idiom - to check if a key is present in the map
	v, ok := m["hello"]
	fmt.Println(v, ok) // 5 true
	v, ok = m["goodbye"]
	fmt.Println(v, ok) //0 false

	// Deleting from Maps
	delete(m, "hello")
	fmt.Println(m)

	fmt.Println(m, len(m))
	clear(m) //Emptying a Map
	fmt.Println(m, len(m))

	m = map[string]int{
		"hello": 5,
		"world": 10,
	}

	n := map[string]int{
		"world": 10,
		"hello": 5,
	}

	fmt.Println(maps.Equal(m, n)) //Comparing Maps

	mapAsSet()
}
