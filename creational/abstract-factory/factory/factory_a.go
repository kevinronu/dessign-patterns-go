package abstractfactory

import (
	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
	branda "github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product/brand-a"
)

type FactoryA struct{}

func (f FactoryA) CreateProductA() product.ProductA {
	return branda.ProductA{}
}

func (f FactoryA) CreateProductB() product.ProductB {
	return branda.ProductB{}
}
