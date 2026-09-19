package product

type ProductType string

type Product interface {
	GetType() ProductType
	DoSomething() string
}
