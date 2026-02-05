package main

import "fmt"

type User struct {
	name          string
	email         string
	paymentMethod string
}

// too much conditionals with complex logic
func (u *User) Payment() {
	if u.paymentMethod == "bank" {
		// logic for bank payment
		fmt.Printf("%s has paid by %s", u.name, u.paymentMethod)
	} else if u.paymentMethod == "paypal" {
		// logic for paypal
	} else if u.paymentMethod == "mobile" {
		// logic for  mobile payment
	} else {
		// logic for apple pay
	}
}

func main() {
	user := User{name: "Rakib", email: "rakib@gmail.com", paymentMethod: "bank"}
	user.Payment()
}
