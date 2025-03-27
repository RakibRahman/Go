package main

import "fmt"

type User struct {
	name  string
	email string
}

// Adhering to SRP. dedicating responsibilies to separate class/method.

// UserService handles database operations - single responsibility

type DatabaseService struct{}

func (db *DatabaseService) SaveUser(user *User) {
	// some database operations
	fmt.Println("Saving user to database", user.name)
}

// EmailService handles mail related operations - single reposibilty

type EmailService struct{}

func (es *EmailService) SendEmail(user *User, message string) {
	// some email sending logic
	fmt.Println(message, user.email)
}

func main() {
	user := User{name: "Rakib", email: "rakib@gmail.com"}

	dbService := &DatabaseService{}
	emailService := &EmailService{}

	dbService.SaveUser(&user)
	emailService.SendEmail(&user, "Welcome to our community")

}
