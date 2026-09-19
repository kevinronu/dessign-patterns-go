package branda

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
)

type ProductB struct{}

func (p ProductB) GetBrand() product.BrandType {
	return BrandA
}

func (p ProductB) DoB() string {
	return fmt.Sprintf("Product B of brand %s is doing B", BrandA)
}
