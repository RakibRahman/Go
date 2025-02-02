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

	var x = []int{1, 2, 3}
	y := []int{20, 30, 40}
	x = append(x, y...) //One slice is appended onto another by using the ... operator to expand the source slice into individual values

	fmt.Println(x)

	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s))
	clear(s)
	fmt.Println(s, len(s))

	fmt.Println("---Slice---")
	str := []string{"a", "b", "c", "d"}
	str2 := str[:2]
	z := str[1:]
	str[1] = "y"
	str2[0] = "x"
	z[1] = "z"
	fmt.Println("str:", str)
	fmt.Println("str2:", str2)
	fmt.Println("z:", z)
	//When you take a slice from a slice, you are not making a copy of the data. Instead,
	// you now have two variables that are sharing memory. This means that changes to an
	// element in a slice affect all slices that share that element.
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
	x := make([]int, 5)
	x = append(x, 10) //the capacity was doubled as soon as the sixth element was appended
	fmt.Println(x, len(x), cap(x))
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
	fmt.Println("---sliceOperation")
	sliceOperation()
	fmt.Println("---sliceUsingMake")
	sliceUsingMake()
	fmt.Println("---iterateOverSlice")
	iterateOverSlice()
}
