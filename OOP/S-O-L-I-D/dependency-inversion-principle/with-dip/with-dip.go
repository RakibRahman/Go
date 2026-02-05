package main

import "fmt"

// Define an interface for sending notifications
type Notifier interface {
	Send(message string)
}

type EmailService struct{}

type SmsService struct{}

// EmailService and SmsService implements the Notifier interface
func (es *EmailService) Send(message string) {
	fmt.Println("Sending Email:", message)
}

func (ss *SmsService) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

// NotificationService depends on the Notifier interface
type NotificationService struct {
	notifier Notifier
}

// Constructor for NotificationService with dependency injection
func NewNotificationService(notifier Notifier) *NotificationService {
	return &NotificationService{
		notifier: notifier,
	}
}

// SendNotification delegates the task to the injected Notifier
func (ns *NotificationService) SendNotification(message string) {
	ns.notifier.Send(message)
}

func withDIP() {
	emailService := &EmailService{}
	smsService := &SmsService{}

	// Inject EmailService into NotificationService
	emailNotificationService := NewNotificationService(emailService)
	emailNotificationService.SendNotification("Keep practicing via email")

	// Inject SmsService into NotificationService
	smsNotificationService := NewNotificationService(smsService)
	smsNotificationService.SendNotification("Keep practicing via SMS")
}

func main() {
	withDIP()
}
