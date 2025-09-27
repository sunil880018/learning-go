// Code should be open for extension but closed for modification. You can achieve this by using interfaces.
package main

import (
	"fmt"
)

// Notification is an interface for different types of notifications.
type Notification interface {
	Send(message string) error
}

type EmailNotification struct{}

func (e *EmailNotification) Send(message string) error {
	fmt.Println("Sending email:", message)
	return nil
}

type SMSNotification struct{}

func (s *SMSNotification) Send(message string) error {
	fmt.Println("Sending SMS:", message)
	return nil
}

func Notify(notification Notification, message string) {
	notification.Send(message)
}

func main() {
	email := &EmailNotification{}
	sms := &SMSNotification{}

	Notify(email, "Welcome to our platform!")
	Notify(sms, "Your verification code is 1234.")
}

// Notification interface allows adding new types (like EmailNotification, SMSNotification) without changing the Notify function.
// The system is open for new notifications but closed for modification.
