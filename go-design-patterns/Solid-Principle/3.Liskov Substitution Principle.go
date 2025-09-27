// A derived class should be able to replace a base class without altering the functionality. In Go, we achieve this through interfaces.
package main

import (
	"fmt"
)

type PaymentProcessor interface {
	Process(amount float64) error
}

type PayPalProcessor struct{}

func (p *PayPalProcessor) Process(amount float64) error {
	fmt.Printf("Processing PayPal payment of %.2f\n", amount)
	return nil
}

type CreditCardProcessor struct{}

func (c *CreditCardProcessor) Process(amount float64) error {
	fmt.Printf("Processing Credit Card payment of %.2f\n", amount)
	return nil
}

func Checkout(p PaymentProcessor, amount float64) {
	p.Process(amount)
}

func main() {
	payPal := &PayPalProcessor{}
	creditCard := &CreditCardProcessor{}

	Checkout(payPal, 100.0)
	Checkout(creditCard, 200.0)
}

// Both PayPalProcessor and CreditCardProcessor implement the PaymentProcessor interface, so they can be
// used interchangeably in the Checkout function.
