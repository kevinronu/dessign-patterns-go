// Package client holds the code that wants the target and knows nothing else.
package client

import "github.com/kevinronu/dessign-patterns-go/structural/adapter/target"

// PaymentService depends only on the target interface, so a new backend never changes this file.
type PaymentService struct {
	Processor target.PaymentProcessor
}

func (s PaymentService) Buy(amountCents int) (string, error) {
	return s.Processor.Pay(amountCents)
}
