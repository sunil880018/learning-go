// Strategy Pattern in Go allows you to dynamically choose an algorithm or behavior at runtime
// without changing the client code. It’s widely used for things like retry logic, payment routing,
// or different calculation strategies.

// 🧩 Scenario

// Suppose you have a payment system that can process payments via Razorpay, Stripe, or PayPal.
// Instead of hardcoding the gateway, you can switch the strategy dynamically.

package main

import "fmt"

// Strategy interface
type PaymentStrategy interface {
	Pay(amount float64)
}

// Concrete Strategy 1 — Razorpay
type RazorpayStrategy struct{}

func (r *RazorpayStrategy) Pay(amount float64) {
	fmt.Printf("Paying ₹%.2f using Razorpay\n", amount)
}

// Concrete Strategy 2 — Stripe
type StripeStrategy struct{}

func (s *StripeStrategy) Pay(amount float64) {
	fmt.Printf("Paying $%.2f using Stripe\n", amount)
}

// Concrete Strategy 3 — PayPal
type PayPalStrategy struct{}

func (p *PayPalStrategy) Pay(amount float64) {
	fmt.Printf("Paying $%.2f using PayPal\n", amount)
}

// Context — uses a payment strategy
type PaymentProcessor struct {
	strategy PaymentStrategy
}

// SetStrategy allows dynamic switching
func (pp *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	pp.strategy = strategy
}

// Execute the strategy
func (pp *PaymentProcessor) Pay(amount float64) {
	pp.strategy.Pay(amount)
}

// Client code
func main() {
	processor := &PaymentProcessor{}

	// Use Razorpay
	processor.SetStrategy(&RazorpayStrategy{})
	processor.Pay(1000)

	// Switch to Stripe dynamically
	processor.SetStrategy(&StripeStrategy{})
	processor.Pay(50)

	// Switch to PayPal
	processor.SetStrategy(&PayPalStrategy{})
	processor.Pay(75)
}

// Explanation

// Strategy interface (PaymentStrategy) → Defines the common method Pay().

// Concrete strategies (RazorpayStrategy, StripeStrategy, PayPalStrategy) → Implement different algorithms.

// Context (PaymentProcessor) → Holds a reference to a strategy and delegates execution.

// Dynamic behavior → Client can switch strategies at runtime without changing the PaymentProcessor code.

// ⚡ Real-world Use Cases

// Payment gateways: Razorpay, Stripe, PayPal routing.

// Retry logic: Exponential backoff, fixed delay, or no retry.

// Compression: Gzip, Brotli, or custom algorithms.

// Sorting: Different sort algorithms based on data type or size.
