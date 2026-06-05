package product

// ProductType represents the type identifier for each concrete product.
type ProductType string

// Product defines the common behavior for all product variants.
type Product interface {
	GetType() ProductType
	DoSomething() string
}
