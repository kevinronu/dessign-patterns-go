// Package target defines the interface the client expects. Every backend has to reach it, either on
// its own or through an adapter.
package target

// PaymentProcessor is the target. Native backends and adapters both sit behind it, so the client
// treats them the same way.
type PaymentProcessor interface {
	Pay(amountCents int) (string, error)
}
