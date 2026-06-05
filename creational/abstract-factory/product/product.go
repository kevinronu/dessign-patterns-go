package product

// BrandType identifies a brand of related products.
type BrandType string

// ProductA is the common contract for the first product kind.
type ProductA interface {
	GetBrand() BrandType
	DoA() string
}

// ProductB is the common contract for the second product kind.
type ProductB interface {
	GetBrand() BrandType
	DoB() string
}
