// Package implementation is the low-level side of the bridge. Each type here knows one output
// format and nothing about documents.
package implementation

// Exporter is the implementation interface: the smallest set of calls a document needs. Keeping it
// small is what makes a new format two short methods.
type Exporter interface {
	Heading(text string) string
	Field(label, value string) string
}
