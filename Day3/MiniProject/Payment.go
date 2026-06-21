package main

import "fmt"

// PaymentProcessor interface (contract)
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

// Concrete implementation: CreditCardProcessor
type CreditCardProcessor struct{}

func (c CreditCardProcessor) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid payment amount: %.2f", amount)
	}
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
	return nil
}
