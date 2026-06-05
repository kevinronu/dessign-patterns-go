package abstractfactory

import (
	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
	branda "github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product/brand-a"
)

// FactoryA creates products of BrandA.
type FactoryA struct{}

// CreateProductA builds a ProductA of BrandA.
func (f FactoryA) CreateProductA() product.ProductA {
	return branda.ProductA{}
}

// CreateProductB builds a ProductB of BrandA.
func (f FactoryA) CreateProductB() product.ProductB {
	return branda.ProductB{}
}
