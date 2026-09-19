package factorymethod

import "github.com/kevinronu/dessign-patterns-go/creational/factory-method/product"

type FactoryA struct{}

func (f FactoryA) CreateProduct(owner string) product.Product {
	return product.ProductA{
		Owner: owner,
		Type:  product.TypeA,
	}
}
