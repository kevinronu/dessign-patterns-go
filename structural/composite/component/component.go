// Package component defines the contract shared by leaves and composites.
package component

// Component is the only type the client needs. A file and a whole tree both implement it, so
// no code has to check which one it holds.
type Component interface {
	Name() string
	Size() int64
	// Find matches the first part of path against its own name, then passes the rest down.
	Find(path string) (Component, bool)
}
