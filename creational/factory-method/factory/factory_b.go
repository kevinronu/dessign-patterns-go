package factorymethod

import "github.com/kevinronu/dessign-patterns-go/creational/factory-method/product"

type FactoryB struct{}

func (f FactoryB) CreateProduct(owner string) product.Product {
	return product.ProductB{
		Owner: owner,
		Type:  product.TypeB,
	}
}
