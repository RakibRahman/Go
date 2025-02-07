package main

import "fmt"

// A unidirectional association where one object contains another, but the contained object can exist independently.

type Professor struct {
	Name string
}

type Department struct {
	Name       string
	Professors []Professor
}

func main() {
	prof1 := Professor{Name: "Dr. Smith"}
	prof2 := Professor{Name: "Dr. Brown"}

	dept := Department{
		Name:       "Science",
		Professors: []Professor{prof1, prof2},
	}

	fmt.Printf("Department: %s\n", dept.Name)
	for _, p := range dept.Professors {
		fmt.Printf("Professor: %s\n", p.Name)
	}

	// Professors can exist independently
	fmt.Println("Independent Professor:", prof1.Name)
}
