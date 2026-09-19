package branda

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
)

type ProductA struct{}

func (p ProductA) GetBrand() product.BrandType {
	return BrandA
}

func (p ProductA) DoA() string {
	return fmt.Sprintf("Product A of brand %s is doing A", BrandA)
}
