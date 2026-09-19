package client

import "github.com/kevinronu/dessign-patterns-go/structural/adapter/target"

type PaymentService struct {
	Processor target.PaymentProcessor
}

func (s PaymentService) Buy(amountCents int) (string, error) {
	return s.Processor.Pay(amountCents)
}
