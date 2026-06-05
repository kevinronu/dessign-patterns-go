package factorymethod

import "github.com/kevinronu/dessign-patterns-go/creational/factory-method/product"

// FactoryB creates ProductB instances.
type FactoryB struct{}

// CreateProduct builds a ProductB with the provided owner.
func (f FactoryB) CreateProduct(owner string) product.Product {
	return product.ProductB{
		Owner: owner,
		Type:  product.TypeB,
	}
}
