// Package adapter connects an adaptee to the target. One adapter per backend that does not fit.
package adapter

import (
	"errors"
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/structural/adapter/adaptee"
)

type LegacyAdapter struct {
	Adaptee adaptee.LegacyService
}

// Pay translates the call and the result between the target and the adaptee's own API.
func (a LegacyAdapter) Pay(amountCents int) (string, error) {
	dollars := float64(amountCents) / 100

	if !a.Adaptee.Charge(dollars) {
		return "", errors.New("legacy service declined the payment")
	}

	return fmt.Sprintf("paid $%.2f via legacy service", dollars), nil
}
