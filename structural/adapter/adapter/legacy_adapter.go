// Package adapter connects an adaptee to the port — one adapter per incompatible backend.
package adapter

import (
	"errors"
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/structural/adapter/adaptee"
)

type LegacyAdapter struct {
	Adaptee adaptee.LegacyService
}

// Pay translates the call and its result between the port and the adaptee's API.
func (a LegacyAdapter) Pay(amountCents int) (string, error) {
	dollars := float64(amountCents) / 100

	if !a.Adaptee.Charge(dollars) {
		return "", errors.New("legacy service declined the payment")
	}

	return fmt.Sprintf("paid $%.2f via legacy service", dollars), nil
}
