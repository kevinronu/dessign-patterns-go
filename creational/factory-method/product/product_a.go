package product

import "fmt"

const (
	// TypeA identifies ProductA.
	TypeA ProductType = "A"
)

// ProductA is a concrete product.
type ProductA struct {
	Owner string
	Type  ProductType
}

// GetType returns the product type.
func (p ProductA) GetType() ProductType {
	return TypeA
}

// DoSomething runs the product's behavior.
func (p ProductA) DoSomething() string {
	return fmt.Sprintf("Product of type %s owned by %s is doing something", p.Type, p.Owner)
}
