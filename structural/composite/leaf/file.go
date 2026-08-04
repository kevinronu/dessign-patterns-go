// Package leaf holds components with no children.
package leaf

import "github.com/kevinronu/dessign-patterns-go/structural/composite/component"

// File is a leaf. Its field is Title and not Name because a struct cannot have a field and a
// method with the same name.
type File struct {
	Title string
	Bytes int64
}

func (f File) Name() string {
	return f.Title
}

func (f File) Size() int64 {
	return f.Bytes
}

// Find can only match itself, because a file has nothing below it.
func (f File) Find(path string) (component.Component, bool) {
	if path != f.Title {
		return nil, false
	}

	return f, true
}
