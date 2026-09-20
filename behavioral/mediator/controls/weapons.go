package controls

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/behavioral/mediator/mediator"
)

type Weapons struct {
	mediator mediator.Mediator

	Enabled bool
	Armed   bool
}

func NewWeapons() *Weapons {
	return &Weapons{Enabled: true}
}

func (w *Weapons) SetMediator(m mediator.Mediator) {
	w.mediator = m
}

func (w *Weapons) Type() mediator.ComponentType {
	return mediator.WeaponsType
}

func (w *Weapons) Enable() {
	w.Enabled = true

	fmt.Println("[Weapons] ENABLED")
}

func (w *Weapons) Disable() {
	w.Enabled = false

	fmt.Println("[Weapons] DISABLED")
}

func (w *Weapons) ToggleArm() {
	if !w.Enabled {
		fmt.Println("[Weapons] Ignored ToggleArm, system disabled")

		return
	}

	w.Armed = !w.Armed

	fmt.Printf("[Weapons] %s\n", armedLabel(w.Armed))

	if w.mediator != nil {
		w.mediator.Notify(w, mediator.WeaponsToggled{Armed: w.Armed})
	}
}

func armedLabel(armed bool) string {
	if armed {
		return "ARMED"
	}

	return "SAFE"
}
