package main

// PaymentService depends on PaymentProcessor
type PaymentService struct {
	processor PaymentProcessor
}

func (s PaymentService) MakePayment(amount float64) error {
	return s.processor.ProcessPayment(amount)
}
