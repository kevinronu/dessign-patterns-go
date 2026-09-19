package partc

import parta "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-a"

type PartC struct {
	monitored *parta.PartA
}

func NewPartC(monitored *parta.PartA) *PartC {
	return &PartC{monitored: monitored}
}

func (p *PartC) ShowStatus() string {
	if p.monitored != nil && p.monitored.IsActive {
		return "monitored PartA is active"
	}

	return "monitored PartA is inactive"
}
