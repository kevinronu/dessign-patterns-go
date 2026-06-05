package branda

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
)

// ProductA is the BrandA implementation of ProductA.
type ProductA struct{}

// GetBrand returns the product's brand.
func (p ProductA) GetBrand() product.BrandType {
	return BrandA
}

// DoA runs the product's behavior.
func (p ProductA) DoA() string {
	return fmt.Sprintf("Product A of brand %s is doing A", BrandA)
}
