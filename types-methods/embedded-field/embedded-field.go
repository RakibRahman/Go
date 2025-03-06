package main

import "fmt"

type Employee struct {
	id   int
	name string
}

type Manager struct {
	Employee
	Reports []Employee
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%d)", e.name, e.id)
}

func (m Manager) FindNewEmployees() Employee {
	// do business logic
	newEmployee := m.Reports[len(m.Reports)-1]
	return newEmployee
}

func main() {
	emp1 := Employee{
		id:   5645,
		name: "rakib",
	}
	emp2 := Employee{
		id:   965,
		name: "zakir",
	}
	emp3 := Employee{
		id:   234,
		name: "nusaiba",
	}
	emp1Desc := emp1.Description()
	fmt.Println(emp1Desc)

	m := Manager{
		Employee: Employee{
			name: "Bob Bobson",
			id:   12345,
		},
		Reports: []Employee{emp1, emp2, emp3},
	}

	fmt.Println(m.Reports)
	fmt.Println(m.FindNewEmployees().name)
}
