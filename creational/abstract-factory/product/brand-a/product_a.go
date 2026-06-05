package branda

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
)

// ProductA is the BrandA implementation of product.ProductA.
type ProductA struct{}

// GetBrand returns the brand for ProductA.
func (p ProductA) GetBrand() product.BrandType {
	return BrandA
}

// DoA executes ProductA behavior for BrandA.
func (p ProductA) DoA() string {
	return fmt.Sprintf("Product A of brand %s is doing A", BrandA)
}
