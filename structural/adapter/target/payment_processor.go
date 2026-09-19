package target

type PaymentProcessor interface {
	Pay(amountCents int) (string, error)
}
