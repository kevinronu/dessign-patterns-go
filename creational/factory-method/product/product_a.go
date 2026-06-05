package product

import "fmt"

const (
	// TypeA identifies ProductA.
	TypeA ProductType = "A"
)

// ProductA is one concrete product implementation.
type ProductA struct {
	Owner string
	Type  ProductType
}

// GetType returns the product type for ProductA.
func (p ProductA) GetType() ProductType {
	return TypeA
}

// DoSomething executes ProductA behavior.
func (p ProductA) DoSomething() string {
	return fmt.Sprintf("Product of type %s owned by %s is doing something", p.Type, p.Owner)
}
