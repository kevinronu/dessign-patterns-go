package subsystem

import (
	"path/filepath"
	"strings"
)

type MediaFile struct {
	Name  string
	Bytes int64
}

// Format reads the codec name from the extension, so no caller has to parse the name itself.
func (f MediaFile) Format() string {
	return strings.TrimPrefix(filepath.Ext(f.Name), ".")
}

// BaseName is the name without the extension, ready to take a new one.
func (f MediaFile) BaseName() string {
	return strings.TrimSuffix(f.Name, filepath.Ext(f.Name))
}
