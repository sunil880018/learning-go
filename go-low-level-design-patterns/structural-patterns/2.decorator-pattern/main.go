// Decorator Pattern in Go allows you to add behavior to objects dynamically without modifying
// their original code. It’s often used for logging, caching, validation, or metrics.

// Scenario

// Suppose you have a basic payment processor and you want to add logging or discount
// features without changing the original code.

package main

import "fmt"

// PaymentProcessor interface
type PaymentProcessor interface {
	Process(amount float64)
}

// Concrete component — Basic payment
type BasicPayment struct{}

func (bp *BasicPayment) Process(amount float64) {
	fmt.Printf("Processing payment of $%.2f\n", amount)
}

// Decorator base — wraps a PaymentProcessor
type PaymentDecorator struct {
	wrapped PaymentProcessor
}

func (pd *PaymentDecorator) Process(amount float64) {
	pd.wrapped.Process(amount)
}

// Concrete decorator — Logging
type LoggingDecorator struct {
	PaymentDecorator
}

func (ld *LoggingDecorator) Process(amount float64) {
	fmt.Println("[LOG] Payment started")
	ld.wrapped.Process(amount)
	fmt.Println("[LOG] Payment finished")
}

// Concrete decorator — Discount
type DiscountDecorator struct {
	PaymentDecorator
	discount float64
}

func (dd *DiscountDecorator) Process(amount float64) {
	amount = amount * (1 - dd.discount)
	dd.wrapped.Process(amount)
}

func main() {
	// Basic payment
	basic := &BasicPayment{}

	// Add logging dynamically
	logged := &LoggingDecorator{PaymentDecorator{wrapped: basic}}
	logged.Process(100)

	fmt.Println("-----")

	// Add discount dynamically
	discounted := &DiscountDecorator{
		PaymentDecorator{wrapped: basic},
		0.1, // 10% discount
	}
	discounted.Process(100)

	fmt.Println("-----")

	// Combine decorators: discount + logging
	combined := &LoggingDecorator{
		PaymentDecorator{wrapped: &DiscountDecorator{
			PaymentDecorator{wrapped: basic},
			0.2, // 20% discount
		}},
	}
	combined.Process(100)
}

// Explanation

// Component (PaymentProcessor) → Interface defining what the system expects.

// Concrete component (BasicPayment) → Original functionality.

// Decorator (PaymentDecorator) → Base wrapper that holds a reference to a component.

// Concrete decorators (LoggingDecorator, DiscountDecorator) → Add extra behavior before/after forwarding the call.

// Dynamic composition → You can chain multiple decorators without modifying the original BasicPayment.

// Real-world use cases

// Adding logging or metrics to services.

// Applying caching or rate-limiting dynamically.

// Enhancing objects with validation, encryption, or discounts.

// Wrapping HTTP handlers in middleware (common in Go web frameworks).
