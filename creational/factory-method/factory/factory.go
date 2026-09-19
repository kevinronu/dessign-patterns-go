package factorymethod

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/factory-method/product"
)

type Factory interface {
	CreateProduct(owner string) product.Product
}

func OneTimeAction(factory Factory, temporaryOwner string) string {
	product := factory.CreateProduct(temporaryOwner)

	return product.DoSomething()
}

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
