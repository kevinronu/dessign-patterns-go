package leaf

import "github.com/kevinronu/dessign-patterns-go/structural/composite/component"

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

func (f File) Find(path string) (component.Component, bool) {
	if path != f.Title {
		return nil, false
	}

	return f, true
}
