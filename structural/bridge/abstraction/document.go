// Package abstraction is the high-level side of the bridge. It decides what a document says and
// leaves the formatting to an implementation.
package abstraction

import "github.com/kevinronu/dessign-patterns-go/structural/bridge/implementation"

// Document is the base abstraction: it holds the implementation so Invoice and Report do not each
// declare it.
type Document struct {
	Exporter implementation.Exporter
}
