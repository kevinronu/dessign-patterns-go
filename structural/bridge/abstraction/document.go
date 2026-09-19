package abstraction

import "github.com/kevinronu/dessign-patterns-go/structural/bridge/implementation"

type Document struct {
	Exporter implementation.Exporter
}
