package composite

import (
	"fmt"
	"slices"
	"strings"

	"github.com/kevinronu/dessign-patterns-go/structural/composite/component"
)

const metadataBytes int64 = 4096

type Folder struct {
	Title    string
	children []component.Component
}

func (f *Folder) Name() string {
	return f.Title
}

func (f *Folder) Add(children ...component.Component) {
	f.children = append(f.children, children...)
}

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
