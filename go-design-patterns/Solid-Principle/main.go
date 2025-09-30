package main

import (
	"fmt"
)

// --- Domain Models ---
type Product struct {
	ID    int
	Name  string
	Price float64
}

type Order struct {
	ID       int
	Products []Product
	Total    float64
}

// --- O: Open-Closed Principle ---
// Payment interface allows extension (new methods) without modifying existing code.
type PaymentProcessor interface {
	Pay(order Order) error
}

// CashPayment implementation
type CashPayment struct{}

func (c *CashPayment) Pay(order Order) error {
	fmt.Println("✅ Paid with Cash")
	return nil
}

// CardPayment implementation
type CardPayment struct{}

func (c *CardPayment) Pay(order Order) error {
	fmt.Println("✅ Paid with Card")
	return nil
}

// UpiPayment implementation
type UpiPayment struct{}

func (u *UpiPayment) Pay(order Order) error {
	fmt.Println("✅ Paid with UPI")
	return nil
}

// --- L: Liskov Substitution Principle ---
// All payment methods can replace each other without breaking behavior.

// --- I: Interface Segregation Principle ---
// Instead of a "fat" interface, we keep PaymentProcessor small & focused.

// --- D: Dependency Inversion Principle ---
// High-level modules depend on abstractions, not concrete classes.

// Notification service interface
type Notifier interface {
	Send(message string) error
}

// Email notifier
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) error {
	fmt.Println("📧 Email sent:", message)
	return nil
}

// SMS notifier
type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) error {
	fmt.Println("📱 SMS sent:", message)
	return nil
}

// --- S: Single Responsibility Principle ---
// OrderService has one job: manage order workflow.
type OrderService struct {
	paymentProcessor PaymentProcessor
	notifier         Notifier
}

func NewOrderService(payment PaymentProcessor, notifier Notifier) *OrderService {
	return &OrderService{
		paymentProcessor: payment,
		notifier:         notifier,
	}
}

func (os *OrderService) Checkout(order Order) error {
	// calculate total
	total := 0.0
	for _, p := range order.Products {
		total += p.Price
	}
	order.Total = total
	fmt.Printf("🛒 Order #%d total: %.2f\n", order.ID, order.Total)

	// process payment
	if err := os.paymentProcessor.Pay(order); err != nil {
		return fmt.Errorf("payment failed: %w", err)
	}

	// send notification
	msg := fmt.Sprintf("Order #%d confirmed. Total: %.2f", order.ID, order.Total)
	if err := os.notifier.Send(msg); err != nil {
		return fmt.Errorf("notification failed: %w", err)
	}

	fmt.Println("🎉 Order completed successfully!")
	return nil
}

// --- Main ---
func main() {
	products := []Product{
		{ID: 1, Name: "Laptop", Price: 75000},
		{ID: 2, Name: "Mouse", Price: 1200},
	}

	order := Order{
		ID:       101,
		Products: products,
	}

	// Example: Pay with Card + Notify via Email
	service := NewOrderService(&CardPayment{}, &EmailNotifier{})
	service.Checkout(order)

	// Example: Pay with UPI + Notify via SMS
	service2 := NewOrderService(&UpiPayment{}, &SMSNotifier{})
	service2.Checkout(order)
}
