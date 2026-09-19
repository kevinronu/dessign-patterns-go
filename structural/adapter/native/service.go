package native

import "fmt"

type Service struct{}

func (Service) Pay(amountCents int) (string, error) {
	return fmt.Sprintf("paid $%.2f via native service", float64(amountCents)/100), nil
}
