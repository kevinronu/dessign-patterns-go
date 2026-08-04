// Package composite holds components that own children.
package composite

import (
	"fmt"
	"slices"
	"strings"

	"github.com/kevinronu/dessign-patterns-go/structural/composite/component"
)

// metadataBytes is what a Folder adds on its own, so a composite is more than the sum of its
// children.
const metadataBytes int64 = 4096

// Folder cannot tell a file from another Folder, so the tree can be as deep as needed.
type Folder struct {
	Title    string
	children []component.Component
}

func (f *Folder) Name() string {
	return f.Title
}

// Add is not part of Component because a leaf has no children to add. Every Folder method takes
// a pointer receiver, so only *Folder implements the contract and the tree cannot hold a copy.
func (f *Folder) Add(children ...component.Component) {
	f.children = append(f.children, children...)
}

// Remove takes a name and not the child itself. Comparing two Component values with == can
// panic when the real type is not comparable, like a struct with a slice inside.
func (f *Folder) Remove(name string) error {
	for i, child := range f.children {
		if child.Name() == name {
			f.children = slices.Delete(f.children, i, i+1)

			return nil
		}
	}

	return fmt.Errorf("folder %q has no child named %q", f.Title, name)
}

func (f *Folder) Size() int64 {
	total := metadataBytes

	for _, child := range f.children {
		total += child.Size()
	}

	return total
}

// Find asks every child for the rest of the path. It does not know which child can answer, and
// it does not need to.
func (f *Folder) Find(path string) (component.Component, bool) {
	head, tail, _ := strings.Cut(path, "/")

	if head != f.Title {
		return nil, false
	}

	if tail == "" {
		return f, true
	}

	for _, child := range f.children {
		if found, ok := child.Find(tail); ok {
			return found, true
		}
	}

	return nil, false
}
