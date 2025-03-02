package main

import "fmt"

//Without Dependency Inversion and Injection

// In this example, the NotificationService is tightly coupled with the EmailService:

type EmailService struct{}
type NotificationService struct {
	emailService *EmailService
}

func (es *EmailService) SendEmail(message string) {
	// perform some task to send email
	fmt.Println("Sending Email:", message)
}

//The NewNotificationService function is a constructor that creates and returns a pointer to a new NotificationService instance.
func NewNotificationService() *NotificationService {
	return &NotificationService{
		emailService: &EmailService{}, // Tight coupling: NotificationService creates its own EmailService
	}
}

func (ns *NotificationService) SendNotification(message string) {
	ns.emailService.SendEmail(message)
}

func withoutDIP() {
	ns := NewNotificationService()
	ns.emailService.SendEmail("Please remember me")
}

func main() {
	withoutDIP()
}
