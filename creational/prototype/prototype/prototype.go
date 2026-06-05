// Package prototype defines the Prototype contract: a value that copies itself.
package prototype

// Prototype is anything that can copy itself. T is the concrete type Clone returns.
type Prototype[T any] interface {
	Clone() T
}
