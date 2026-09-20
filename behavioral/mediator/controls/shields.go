package controls

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/behavioral/mediator/mediator"
)

type Shields struct {
	mediator mediator.Mediator

	Enabled bool
	Up      bool
}

func NewShields() *Shields {
	return &Shields{Enabled: true}
}

func (s *Shields) SetMediator(m mediator.Mediator) {
	s.mediator = m
}

func (s *Shields) Type() mediator.ComponentType {
	return mediator.ShieldsType
}

func (s *Shields) Enable() {
	s.Enabled = true

	fmt.Println("[Shields] ENABLED")
}

func (s *Shields) Disable() {
	s.Enabled = false

	fmt.Println("[Shields] DISABLED")
}

func (s *Shields) ToggleShields() {
	if !s.Enabled {
		fmt.Println("[Shields] Ignored ToggleShields, system disabled")

		return
	}

	s.Up = !s.Up

	fmt.Printf("[Shields] %s\n", upLabel(s.Up))

	if s.mediator != nil {
		s.mediator.Notify(s, mediator.ShieldsToggled{Up: s.Up})
	}
}

func upLabel(up bool) string {
	if up {
		return "RAISED"
	}

	return "LOWERED"
}
