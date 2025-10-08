// 🧩 Scenario

// You have a payment interface that your application uses, but a new third-party service (like Razorpay)
// provides a different API.
// Instead of changing your entire codebase, you write an adapter to make it compatible.

// 🚀 Real-world Use Cases

// Integrating multiple payment gateways (Razorpay, Stripe, PayPal).

// Logging systems (adapting from zap to logrus).

// Using different caching or messaging systems (Redis, Memcached, Kafka clients).

package main

import "fmt"

// Target interface — what your app expects
type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

// Adaptee 1 — Razorpay (third-party library)
type Razorpay struct{}

// Adaptee 2 — Stripe (another third-party library)
type Stripe struct{}

// Adapter 1 — RazorpayAdapter makes Razorpay compatible with PaymentProcessor
type RazorpayAdapter struct {
	gateway *Razorpay
}

// Adapter 2 — StripeAdapter makes Stripe compatible with PaymentProcessor
type StripeAdapter struct {
	gateway *Stripe
}

// Razorpay’s own method (incompatible with PaymentProcessor)
func (r *Razorpay) MakePayment(amount float64) {
	fmt.Printf("Processing payment of ₹%.2f using Razorpay API\n", amount)
}

// Stripe’s own method (incompatible with PaymentProcessor)
func (s *Stripe) Pay(amountInCents int64) {
	fmt.Printf("Processing payment of $%.2f using Stripe API\n", float64(amountInCents)/100)
}

// RazorpayAdapter implements PaymentProcessor interface
func (ra *RazorpayAdapter) ProcessPayment(amount float64) {
	ra.gateway.MakePayment(amount)
}

// StripeAdapter implements PaymentProcessor interface
func (sa *StripeAdapter) ProcessPayment(amount float64) {
	cents := int64(amount * 100)
	sa.gateway.Pay(cents)
}

// Client code
func main() {
	var processor PaymentProcessor

	// Use Razorpay via adapter
	processor = &RazorpayAdapter{gateway: &Razorpay{}}
	processor.ProcessPayment(999.99)

	// Use Stripe via adapter
	processor = &StripeAdapter{gateway: &Stripe{}}
	processor.ProcessPayment(49.99)
}

// Explanation

// Both Razorpay and Stripe have different APIs (MakePayment vs Pay).

// We define a common interface PaymentProcessor so the rest of the system doesn’t care about implementation differences.

// Adapters (RazorpayAdapter, StripeAdapter) translate calls into the appropriate format.

// The client (main) simply uses the PaymentProcessor interface — achieving loose coupling and high flexibility.
