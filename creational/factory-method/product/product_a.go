package product

import "fmt"

const (
	TypeA ProductType = "A"
)

type ProductA struct {
	Owner string
	Type  ProductType
}

func (p ProductA) GetType() ProductType {
	return TypeA
}

func (p ProductA) DoSomething() string {
	return fmt.Sprintf("Product of type %s owned by %s is doing something", p.Type, p.Owner)
}
