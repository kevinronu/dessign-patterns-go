package brandb

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
)

type ProductA struct{}

func (p ProductA) GetBrand() product.BrandType {
	return BrandB
}

func (p ProductA) DoA() string {
	return fmt.Sprintf("Product A of brand %s is doing A", BrandB)
}
