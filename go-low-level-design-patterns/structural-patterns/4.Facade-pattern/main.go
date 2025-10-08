// The Facade Design Pattern in Go provides a simplified interface to a complex subsystem.
// It hides the complexity and allows the client to interact with the system in a simple and unified way.

// This is often used in payment systems, messaging services, or any system with multiple subsystems.

// Scenario

// Suppose you have a payment system with multiple steps:

// 1.Validate the user.

// 2.Process the payment.

// 3.Send notification.

// Instead of the client calling each subsystem individually, we create a Facade to simplify it.

package main

import "fmt"

// Subsystem 1: User validation
type UserValidator struct{}

func (uv *UserValidator) ValidateUser(userID string) bool {
	fmt.Println("Validating user:", userID)
	return true
}

// Subsystem 2: Payment processor
type PaymentProcessor struct{}

func (pp *PaymentProcessor) ProcessPayment(userID string, amount float64) {
	fmt.Printf("Processing payment of $%.2f for user %s\n", amount, userID)
}

// Subsystem 3: Notification service
type Notifier struct{}

func (n *Notifier) SendNotification(userID string, message string) {
	fmt.Printf("Sending notification to %s: %s\n", userID, message)
}

// Facade
type PaymentFacade struct {
	validator *UserValidator
	processor *PaymentProcessor
	notifier  *Notifier
}

func NewPaymentFacade() *PaymentFacade {
	return &PaymentFacade{
		validator: &UserValidator{},
		processor: &PaymentProcessor{},
		notifier:  &Notifier{},
	}
}

func (pf *PaymentFacade) MakePayment(userID string, amount float64) {
	if pf.validator.ValidateUser(userID) {
		pf.processor.ProcessPayment(userID, amount)
		pf.notifier.SendNotification(userID, "Payment successful!")
	}
}

// Client code
func main() {
	paymentSystem := NewPaymentFacade()
	paymentSystem.MakePayment("user123", 250.0)
}

// Explanation

// Subsystems → UserValidator, PaymentProcessor, Notifier are complex and independent.

// Facade (PaymentFacade) → Provides a single method MakePayment() to the client.

// Client → Only interacts with the facade, unaware of the underlying subsystems.
