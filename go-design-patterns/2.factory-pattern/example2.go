package main

import (
    "fmt"
)

type PaymentMethod interface {
    Pay(amount float64)
}

// CreditCard pay via credit card
type CreditCard struct{}
func (c *CreditCard) Pay(amount float64) {
    fmt.Printf("Paying %.2f using credit card.\\n", amount)
}

// DebitCard pay via debit card
type DebitCard struct{}
func (d *DebitCard) Pay(amount float64) {
    fmt.Printf("Paying %.2f using debit card.\\n", amount)
}

// DebitCard pay via cash
type Cash struct{}
func (c *Cash) Pay(amount float64) {
    fmt.Printf("Paying %.2f using cash.\\n", amount)
}

type PaymentMethodFactory struct{}

func (pmf *PaymentMethodFactory) CreatePaymentMethod(paymentType string) (PaymentMethod, error) {
    switch paymentType {
    case "credit":
        return &CreditCard{}, nil
    case "debit":
        return &DebitCard{}, nil
    case "cash":
        return &Cash{}, nil
    default:
        return nil, fmt.Errorf("Invalid payment method type: %s", paymentType)
    }
}

func main() {
    paymentMethodFactory := &PaymentMethodFactory{}

    paymentMethod, err := paymentMethodFactory.CreatePaymentMethod("credit")
    if err != nil {
        fmt.Println(err)
        return
    }
    paymentMethod.Pay(100.0)

    paymentMethod, err = paymentMethodFactory.CreatePaymentMethod("debit")
    if err != nil {
        fmt.Println(err)
        return
    }
    paymentMethod.Pay(50.0)

    paymentMethod, err = paymentMethodFactory.CreatePaymentMethod("cash")
    if err != nil {
        fmt.Println(err)
        return
    }
    paymentMethod.Pay(20.0)
}


// 2. Factory Pattern in Go

// ✅ Concept

// 	Create objects without exposing the creation logic to the client.
// 	Useful when object creation depends on conditions, config, or environment.

// ✅ Real-World Go Use Cases

// 	Database driver selection

// 	You want to support multiple DBs (Postgres, MySQL, SQLite). The factory decides which to initialize.

// 	Logger factory

// 	Create different loggers (file logger, console logger, JSON logger).

// 	Message queue producers/consumers

// 	Kafka vs RabbitMQ vs SQS → factory picks correct client.

// 	Payment gateways

// 	Razorpay, Stripe, PayPal → factory returns correct gateway implementation.

// 	Cloud providers SDK

// 	AWS S3 vs GCP Storage vs Azure Blob → factory builds correct client.