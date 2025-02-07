package main

import "fmt"

//Two objects exist independently but can interact.

type Student struct {
	Name string
}

type Teacher struct {
	Name string
}

func (t *Teacher) Teach(s *Student) {
	fmt.Printf("%s is teaching %s", t.Name, s.Name)
}

func main() {
	teacher := Teacher{Name: "Zakir Hossain"}
	student := Student{Name: "Rakib"}
	teacher.Teach(&student) //The Teacher and Student structs are independent but interact through the Teach() method.
}
