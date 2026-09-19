package extrinsic

import "github.com/kevinronu/dessign-patterns-go/structural/flyweight/flyweight"

type Tree struct {
	X    int
	Y    int
	Type *flyweight.TreeType
}

func (t Tree) Draw() string {
	return t.Type.Draw(t.X, t.Y)
}
