package factorymethod

import "github.com/kevinronu/dessign-patterns-go/creational/factory-method/product"

// FactoryA creates a ProductA.
type FactoryA struct{}

// CreateProduct builds a ProductA for the given owner.
func (f FactoryA) CreateProduct(owner string) product.Product {
	return product.ProductA{
		Owner: owner,
		Type:  product.TypeA,
	}
}
