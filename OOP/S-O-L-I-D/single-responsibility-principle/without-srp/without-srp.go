package main

import "fmt"

type User struct {
	name  string
	email string
}

// Violation of SRP. its handling both database operations and email sending.

func (u *User) Save() error {
	// some logic to save users on database
	fmt.Printf("Saving user %s to database\n", u.name)

	// some logic to send email
	u.SendEmail(u.email, "Welcome to our community")
	return nil
}

func (u *User) SendEmail(email, message string) {
	// Email sending logic
	fmt.Println(message, email)
}

func main() {
	user := User{name: "Rakib", email: "rakib@gmail.com"}
	user.Save()
}
