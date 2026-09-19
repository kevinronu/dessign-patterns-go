package component

type Component interface {
	Name() string
	Size() int64
	Find(path string) (Component, bool)
}
