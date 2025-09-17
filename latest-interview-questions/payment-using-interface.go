package main

import "fmt"

// PaymentGateway defines how a payment system should behave
type PaymentGateway interface {
	Pay(amount float64) error
	Refund(transactionID string) error
}

// Razorpay implements PaymentGateway
type Razorpay struct{}

func (r Razorpay) Pay(amount float64) error {
	fmt.Printf("Processing payment of ₹%.2f via Razorpay...\n", amount)
	return nil
}

func (r Razorpay) Refund(transactionID string) error {
	fmt.Printf("Refunding transaction %s via Razorpay...\n", transactionID)
	return nil
}

// Stripe implements PaymentGateway
type Stripe struct{}

func (s Stripe) Pay(amount float64) error {
	fmt.Printf("Processing payment of $%.2f via Stripe...\n", amount)
	return nil
}

func (s Stripe) Refund(transactionID string) error {
	fmt.Printf("Refunding transaction %s via Stripe...\n", transactionID)
	return nil
}

// Customer uses any PaymentGateway
type Customer struct {
	Name    string
	Gateway PaymentGateway
}

func (c Customer) MakePayment(amount float64) {
	fmt.Println("Name : ", c.Name)
	fmt.Println("Payment Gateway : ", c.Gateway)
	err := c.Gateway.Pay(amount)
	if err != nil {
		fmt.Println("Payment failed:", err)
	} else {
		fmt.Println("Payment successful for", c.Name)
	}
}

func (c Customer) RequestRefund(txnID string) {
	err := c.Gateway.Refund(txnID)
	if err != nil {
		fmt.Println("Refund failed:", err)
	} else {
		fmt.Println("Refund successful for", c.Name)
	}
}

func main() {
	// Customer using Razorpay
	customer1 := Customer{Name: "Ramesh", Gateway: Razorpay{}}
	customer1.MakePayment(1500.50)
	customer1.RequestRefund("TXN12345")

	// Customer using Stripe
	customer2 := Customer{Name: "Sunil", Gateway: Stripe{}}
	customer2.MakePayment(250.75)
	customer2.RequestRefund("TXN98765")
}
