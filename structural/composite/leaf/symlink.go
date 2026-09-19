package leaf

import "github.com/kevinronu/dessign-patterns-go/structural/composite/component"

type Symlink struct {
	Title  string
	Target string
}

func (s Symlink) Name() string {
	return s.Title
}

func (s Symlink) Size() int64 {
	return int64(len(s.Target))
}

func (s Symlink) Find(path string) (component.Component, bool) {
	if path != s.Title {
		return nil, false
	}

	return s, true
}
