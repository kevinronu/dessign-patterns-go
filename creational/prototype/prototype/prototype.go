package prototype

type Prototype[T any] interface {
	Clone() T
}
