package main

import "fmt"

func main() {
	// Choose implementation
	processor := CreditCardProcessor{}

	// Inject into service
	service := PaymentService{processor: processor}

	// Try a valid payment
	err := service.MakePayment(100.0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Try an invalid payment
	err = service.MakePayment(-50.0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
