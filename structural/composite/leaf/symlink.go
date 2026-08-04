package leaf

import "github.com/kevinronu/dessign-patterns-go/structural/composite/component"

// Symlink calculates its size, which is why the contract asks for a method and not a field.
type Symlink struct {
	Title  string
	Target string
}

func (s Symlink) Name() string {
	return s.Title
}

// On Unix the size of a symlink is the length of its target path.
func (s Symlink) Size() int64 {
	return int64(len(s.Target))
}

// Find does not follow the target. Following it needs the whole tree, and a leaf only knows
// itself.
func (s Symlink) Find(path string) (component.Component, bool) {
	if path != s.Title {
		return nil, false
	}

	return s, true
}
