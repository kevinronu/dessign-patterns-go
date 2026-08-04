// Package extrinsic holds the context object. GoF calls this role Context, but in Go that name
// belongs to the standard library, so the package is named after the state it carries.
package extrinsic

import "github.com/kevinronu/dessign-patterns-go/structural/flyweight/flyweight"

// Tree is the context: what differs for every tree, plus a pointer to what does not. Millions of
// these exist, so it holds nothing else.
type Tree struct {
	X    int
	Y    int
	Type *flyweight.TreeType
}

func (t Tree) Draw() string {
	return t.Type.Draw(t.X, t.Y)
}
