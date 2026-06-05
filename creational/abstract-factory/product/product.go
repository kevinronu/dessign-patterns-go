package product

// BrandType represents the brand identifier shared by related products.
type BrandType string

// ProductA defines the common behavior for the first product kind.
type ProductA interface {
	GetBrand() BrandType
	DoA() string
}

// ProductB defines the common behavior for the second product kind.
type ProductB interface {
	GetBrand() BrandType
	DoB() string
}
