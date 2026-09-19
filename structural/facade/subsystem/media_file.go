package subsystem

import (
	"path/filepath"
	"strings"
)

type MediaFile struct {
	Name  string
	Bytes int64
}

func (f MediaFile) Format() string {
	return strings.TrimPrefix(filepath.Ext(f.Name), ".")
}

func (f MediaFile) BaseName() string {
	return strings.TrimSuffix(f.Name, filepath.Ext(f.Name))
}
