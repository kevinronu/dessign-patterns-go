// Package port defines the PaymentProcessor port: the plug-in point the client owns and every backend connects to.
package port

// PaymentProcessor is the port. The client calls it; native backends and adapters plug in behind it, so the client treats them all the same.
type PaymentProcessor interface {
	Pay(amountCents int) (string, error)
}
