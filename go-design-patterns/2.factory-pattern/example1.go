package main

import "fmt"

// Interface
type PaymentGateway interface {
	Pay(amount float64) string
}

// Implementations
type Razorpay struct{}

func (r Razorpay) Pay(amount float64) string { return fmt.Sprintf("Paid %.2f via Razorpay", amount) }

type Stripe struct{}

func (s Stripe) Pay(amount float64) string { return fmt.Sprintf("Paid %.2f via Stripe", amount) }

// Factory
func GetPaymentGateway(provider string) PaymentGateway {
	switch provider {
	case "razorpay":
		return Razorpay{}
	case "stripe":
		return Stripe{}
	default:
		return nil
	}
}

func main() {
	payment1 := GetPaymentGateway("razorpay")
	fmt.Println(payment1.Pay(100.0))
}
