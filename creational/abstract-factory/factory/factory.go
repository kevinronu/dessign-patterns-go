package abstractfactory

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
	branda "github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product/brand-a"
	brandb "github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product/brand-b"
)

type Factory interface {
	CreateProductA() product.ProductA
	CreateProductB() product.ProductB
}

func OneTimeAction(factory Factory) string {
	productA := factory.CreateProductA()
	productB := factory.CreateProductB()

	return fmt.Sprintf("%s. %s", productA.DoA(), productB.DoB())
}

func GetFactory(brand product.BrandType) (Factory, error) {
	switch brand {
	case branda.BrandA:
		return FactoryA{}, nil
	case brandb.BrandB:
		return FactoryB{}, nil
	default:
		return nil, fmt.Errorf("unsupported product brand: %q", brand)
	}
}
