package controls

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/behavioral/mediator/mediator"
)

type Thruster struct {
	mediator mediator.Mediator

	Enabled bool
	Thrust  int
}

func NewThruster() *Thruster {
	return &Thruster{Enabled: true}
}

func (t *Thruster) SetMediator(m mediator.Mediator) {
	t.mediator = m
}

func (t *Thruster) Type() mediator.ComponentType {
	return mediator.ThrusterType
}

func (t *Thruster) Enable() {
	t.Enabled = true

	fmt.Println("[Thrusters] ENABLED")
}

func (t *Thruster) Disable() {
	t.Enabled = false

	fmt.Println("[Thrusters] DISABLED")
}

func (t *Thruster) SetThrust(percent int) {
	if !t.Enabled {
		fmt.Println("[Thrusters] Ignored SetThrust, system disabled")

		return
	}

	t.Thrust = min(100, max(0, percent))

	fmt.Printf("[Thrusters] thrust -> %d%%\n", t.Thrust)

	if t.mediator != nil {
		t.mediator.Notify(t, mediator.ThrustSet{Percent: t.Thrust})
	}
}
