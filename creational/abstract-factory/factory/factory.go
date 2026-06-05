package abstractfactory

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
	branda "github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product/brand-a"
	brandb "github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product/brand-b"
)

// Factory defines the behavior required to create a related brand of products.
type Factory interface {
	CreateProductA() product.ProductA
	CreateProductB() product.ProductB
}

// OneTimeAction creates a full brand of products and executes their behavior.
func OneTimeAction(factory Factory) string {
	productA := factory.CreateProductA()
	productB := factory.CreateProductB()

	return fmt.Sprintf("%s. %s", productA.DoA(), productB.DoB())
}

// GetFactory returns a factory implementation for the requested brand type.
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
