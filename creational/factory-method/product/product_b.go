package product

import "fmt"

const (
	TypeB ProductType = "B"
)

type ProductB struct {
	Owner string
	Type  ProductType
}

func (p ProductB) GetType() ProductType {
	return TypeB
}

func (p ProductB) DoSomething() string {
	return fmt.Sprintf("Product of type %s owned by %s is doing something", p.Type, p.Owner)
}
