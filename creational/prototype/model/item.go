// Package model holds concrete prototypes that implement prototype.Prototype.
package model

// Item is a leaf prototype: it holds no nested references.
type Item struct {
	Name string
	Tag  string
}

// Clone returns a copy of the Item.
func (i Item) Clone() Item {
	return Item{Name: i.Name + "_clone", Tag: i.Tag}
}
