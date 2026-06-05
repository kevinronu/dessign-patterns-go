package product

// ProductType identifies a concrete product.
type ProductType string

// Product is the common contract for every concrete product.
type Product interface {
	GetType() ProductType
	DoSomething() string
}
