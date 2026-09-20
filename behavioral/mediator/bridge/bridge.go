package bridge

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/behavioral/mediator/controls"
	"github.com/kevinronu/dessign-patterns-go/behavioral/mediator/mediator"
)

const safeThrust = 60

type Bridge struct {
	thruster *controls.Thruster
	weapons  *controls.Weapons
	shields  *controls.Shields
	scanner  *controls.Scanner
}

func New(components ...mediator.Component) *Bridge {
	bridge := &Bridge{}

	for _, component := range components {
		bridge.Register(component)
	}

	return bridge
}

func (b *Bridge) Register(component mediator.Component) {
	component.SetMediator(b)

	switch control := component.(type) {
	case *controls.Thruster:
		b.thruster = control
	case *controls.Weapons:
		b.weapons = control
	case *controls.Shields:
		b.shields = control
	case *controls.Scanner:
		b.scanner = control
	}
}

func (b *Bridge) Notify(sender mediator.Component, event mediator.Event) {
	if !mediator.Allows(sender.Type(), event.Name()) {
		fmt.Printf("\n[Bridge] rejected %s from %s\n", event.Name(), sender.Type())

		return
	}

	fmt.Printf("\n[Bridge] %s from %s\n", event.Name(), sender.Type())

	switch event := event.(type) {
	case mediator.ThrustSet:
		b.dropShieldsAtHighThrust(event.Percent)
		b.capThrustWhileArmed(event.Percent)
	case mediator.WeaponsToggled:
		if event.Armed {
			b.capThrustWhileArmed(b.thruster.Thrust)
		}
	case mediator.ShieldsToggled:
		b.followShields(event.Up)
	case mediator.ScanRequested:
		b.runScan()
	}

	b.reportStatus()
}

func (b *Bridge) dropShieldsAtHighThrust(thrust int) {
	if thrust <= 80 {
		return
	}

	fmt.Println("[Bridge] Thrust > 80% -> forcing shields DOWN")
	b.shields.Up = false
	b.shields.Disable()
}

// capThrustWhileArmed writes the thrust field instead of calling SetThrust,
// because that action would notify the bridge again.
func (b *Bridge) capThrustWhileArmed(thrust int) {
	if !b.weapons.Armed || thrust <= safeThrust {
		return
	}

	fmt.Printf("[Bridge] Weapons armed -> capping thrust at %d%%\n", safeThrust)
	b.thruster.Thrust = safeThrust
}

func (b *Bridge) followShields(up bool) {
	if !up {
		fmt.Println("[Bridge] Shields lowered -> auto-initiating scanner")
		b.runScan()

		return
	}

	if b.weapons.Armed {
		fmt.Println("[Bridge] Shields UP -> auto-disarming weapons")
		b.weapons.Armed = false
	}
}

func (b *Bridge) runScan() {
	if !b.scanner.Enabled {
		fmt.Println("[Bridge] Scanner disabled -> scan skipped")

		return
	}

	fmt.Println("[Bridge] Scan in progress -> disabling thrusters")
	b.thruster.Disable()

	b.scanner.Run()

	fmt.Println("[Bridge] Scan complete -> re-enabling thrusters")
	b.thruster.Enable()
}

func (b *Bridge) reportStatus() {
	fmt.Printf("[Bridge::Status] Thrust=%d%% | Weapons=%v | Shields=%v | LastScan=%q\n",
		b.thruster.Thrust, b.weapons.Armed, b.shields.Up, b.scanner.LastScan)
}
