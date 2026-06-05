package product

// Product is what a builder produces. Each concrete product renders itself differently.
type Product interface {
	Describe() string
}
