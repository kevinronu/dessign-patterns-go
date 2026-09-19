package abstractfactory

import (
	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
	brandb "github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product/brand-b"
)

type FactoryB struct{}

func (f FactoryB) CreateProductA() product.ProductA {
	return brandb.ProductA{}
}

func (f FactoryB) CreateProductB() product.ProductB {
	return brandb.ProductB{}
}
