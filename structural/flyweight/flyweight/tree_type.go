// Package flyweight holds the shared object and the factory that hands it out.
package flyweight

import "fmt"

// TreeType is the flyweight: the state every tree of one kind shares. Its fields are private, so
// one can only come from the factory, and that is what keeps them shared.
type TreeType struct {
	name    string
	color   string
	texture string
}

// Draw takes the position as arguments instead of storing it. That is what lets one TreeType work
// for every tree of its kind.
func (t *TreeType) Draw(x, y int) string {
	return fmt.Sprintf("%s at (%d,%d) in %s", t.name, x, y, t.color)
}
