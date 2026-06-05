package brandb

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
)

// ProductB is the BrandB implementation of ProductB.
type ProductB struct{}

// GetBrand returns the product's brand.
func (p ProductB) GetBrand() product.BrandType {
	return BrandB
}

// DoB runs the product's behavior.
func (p ProductB) DoB() string {
	return fmt.Sprintf("Product B of brand %s is doing B", BrandB)
}
