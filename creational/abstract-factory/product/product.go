package product

type BrandType string

type ProductA interface {
	GetBrand() BrandType
	DoA() string
}

type ProductB interface {
	GetBrand() BrandType
	DoB() string
}
