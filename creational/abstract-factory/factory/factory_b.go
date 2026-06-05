package abstractfactory

import (
	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
	brandb "github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product/brand-b"
)

// FactoryB creates products of BrandB.
type FactoryB struct{}

// CreateProductA builds a ProductA of BrandB.
func (f FactoryB) CreateProductA() product.ProductA {
	return brandb.ProductA{}
}

// CreateProductB builds a ProductB of BrandB.
func (f FactoryB) CreateProductB() product.ProductB {
	return brandb.ProductB{}
}
