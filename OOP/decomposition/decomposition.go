package main

import "fmt"

//  Breaking down a system into smaller, manageable objects.

type Student struct {
	Name string
}

type Professor struct {
	Name string
}

type Course struct {
	Title     string
	Professor Professor
}

type Department struct {
	Name    string
	Courses []Course
}

func main() {
	prof := Professor{Name: "Saleh"}
	course := Course{Title: "Mathematics", Professor: prof}
	course2 := Course{Title: "Physics", Professor: prof}

	dept := Department{Name: "Science", Courses: []Course{course, course2}}
	fmt.Printf("Department: %s\n", dept.Name)
	for _, c := range dept.Courses {
		fmt.Printf("Course: %s, Taught by: %s\n", c.Title, c.Professor.Name)
	}
}
