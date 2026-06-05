package branda

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
)

// ProductB is the BrandA implementation of product.ProductB.
type ProductB struct{}

// GetBrand returns the brand for ProductB.
func (p ProductB) GetBrand() product.BrandType {
	return BrandA
}

// DoB executes ProductB behavior for BrandA.
func (p ProductB) DoB() string {
	return fmt.Sprintf("Product B of brand %s is doing B", BrandA)
}
