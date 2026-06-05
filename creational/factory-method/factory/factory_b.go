package factorymethod

import "github.com/kevinronu/dessign-patterns-go/creational/factory-method/product"

// FactoryB creates a ProductB.
type FactoryB struct{}

// CreateProduct builds a ProductB for the given owner.
func (f FactoryB) CreateProduct(owner string) product.Product {
	return product.ProductB{
		Owner: owner,
		Type:  product.TypeB,
	}
}
