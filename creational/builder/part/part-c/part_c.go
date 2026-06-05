package partc

import parta "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-a"

// PartC is an optional part that observes a PartA and reports its state.
type PartC struct {
	monitored *parta.PartA
}

// NewPartC binds a PartC to the PartA it observes.
func NewPartC(monitored *parta.PartA) *PartC {
	return &PartC{monitored: monitored}
}

// ShowStatus describes the observed PartA's state.
func (p *PartC) ShowStatus() string {
	if p.monitored != nil && p.monitored.IsActive {
		return "monitored PartA is active"
	}

	return "monitored PartA is inactive"
}
