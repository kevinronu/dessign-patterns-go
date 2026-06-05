package product

import "fmt"

const (
	// TypeB identifies ProductB.
	TypeB ProductType = "B"
)

// ProductB is one concrete product implementation.
type ProductB struct {
	Owner string
	Type  ProductType
}

// GetType returns the product type for ProductB.
func (p ProductB) GetType() ProductType {
	return TypeB
}

// DoSomething executes ProductB behavior.
func (p ProductB) DoSomething() string {
	return fmt.Sprintf("Product of type %s owned by %s is doing something", p.Type, p.Owner)
}
