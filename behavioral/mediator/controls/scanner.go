package controls

import (
	"fmt"
	"math/rand/v2"

	"github.com/kevinronu/dessign-patterns-go/behavioral/mediator/mediator"
)

var readings = []string{"asteroid", "nebula", "hostiles", "clear"}

type Scanner struct {
	mediator mediator.Mediator

	Enabled  bool
	LastScan string
}

func NewScanner() *Scanner {
	return &Scanner{Enabled: true, LastScan: "none"}
}

func (s *Scanner) SetMediator(m mediator.Mediator) {
	s.mediator = m
}

func (s *Scanner) Type() mediator.ComponentType {
	return mediator.ScannerType
}

func (s *Scanner) Enable() {
	s.Enabled = true

	fmt.Println("[Scanner] ENABLED")
}

func (s *Scanner) Disable() {
	s.Enabled = false

	fmt.Println("[Scanner] DISABLED")
}

func (s *Scanner) RequestScan() {
	if !s.Enabled {
		fmt.Println("[Scanner] Ignored RequestScan, system disabled")

		return
	}

	fmt.Println("[Scanner] scan requested")

	if s.mediator != nil {
		s.mediator.Notify(s, mediator.ScanRequested{})
	}
}

func (s *Scanner) Run() {
	s.LastScan = readings[rand.IntN(len(readings))]

	fmt.Printf("[Scanner] scan result: %q\n", s.LastScan)
}
