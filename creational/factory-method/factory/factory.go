package factorymethod

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/factory-method/product"
)

// Factory creates a product.
type Factory interface {
	CreateProduct(owner string) product.Product
}

// OneTimeAction creates a temporary product and runs its behavior.
func OneTimeAction(factory Factory, temporaryOwner string) string {
	product := factory.CreateProduct(temporaryOwner)

	return product.DoSomething()
}

// GetFactory returns the factory for the given product type.
func GetFactory(productType product.ProductType) (Factory, error) {
	switch productType {
	case product.TypeA:
		return FactoryA{}, nil
	case product.TypeB:
		return FactoryB{}, nil
	default:
		return nil, fmt.Errorf("unsupported product type: %q", productType)
	}
}
