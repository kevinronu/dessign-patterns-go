// Package implementation is the bridge's low-level side: the dependency the abstraction
// builds on. Each concrete type provides the formatting primitives for one format.
package implementation

// Exporter is the implementation interface: the primitives the abstraction composes a
// document from. Each format renders them its own way.
type Exporter interface {
	Heading(text string) string
	Field(label, value string) string
}
