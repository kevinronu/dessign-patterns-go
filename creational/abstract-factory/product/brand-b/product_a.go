package brandb

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
)

// ProductA is the BrandB implementation of product.ProductA.
type ProductA struct{}

// GetBrand returns the brand for ProductA.
func (p ProductA) GetBrand() product.BrandType {
	return BrandB
}

// DoA executes ProductA behavior for BrandB.
func (p ProductA) DoA() string {
	return fmt.Sprintf("Product A of brand %s is doing A", BrandB)
}
