// Package abstraction is the bridge's high-level side: the policy that depends on an
// implementation and composes a document from its primitives.
package abstraction

import "github.com/kevinronu/dessign-patterns-go/structural/bridge/implementation"

// Document is the base abstraction: it holds the implementation the refined documents build their output with.
type Document struct {
	Exporter implementation.Exporter
}
