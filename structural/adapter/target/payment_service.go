// Package target holds the client that depends on the port.
package target

import "github.com/kevinronu/dessign-patterns-go/structural/adapter/port"

// PaymentService is the client: it depends only on the PaymentProcessor port, so any
// backend plugs in without changing it.
type PaymentService struct {
	Processor port.PaymentProcessor
}

func (s PaymentService) Buy(amountCents int) (string, error) {
	return s.Processor.Pay(amountCents)
}
