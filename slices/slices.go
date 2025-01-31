package main

import (
	"fmt"
	"slices"
)

func sliceOperation() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println(slice)

	//delete element
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println(slice)
	slice[0] = 99
	fmt.Println(slice)
	// fmt.Println(slice[1:3])
}

func iterateOverSlice() {
	slice := []int{10, 20, 30}

	//using index
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//using range
	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}

func sliceUsingMake() {
	makeSlice := make([]int, 5) // Creates a slice of length 5, capacity 5
	makeSlice2 := make([]string, 3, 5)
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))
	fmt.Println("---")
	fmt.Println(len(makeSlice2))
	fmt.Println(cap(makeSlice2))
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	var sparseSlice = []int{1, 2, 3: 4, 5, 6: 10}
	s := []string{"a", "b", "c"}
	fmt.Println(x)
	sparseSlice[3] = 56
	fmt.Println(sparseSlice)
	fmt.Println(sparseSlice == nil)
	fmt.Println(len(s))
	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	fmt.Println("---sliceCRUD")
	sliceOperation()
	fmt.Println("---sliceOperation")
	sliceUsingMake()
	fmt.Println("---iterateOverSlice")
	iterateOverSlice()
}
