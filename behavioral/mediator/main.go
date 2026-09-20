package main

import (
	"github.com/kevinronu/dessign-patterns-go/behavioral/mediator/bridge"
	"github.com/kevinronu/dessign-patterns-go/behavioral/mediator/controls"
)

func main() {
	thruster := controls.NewThruster()
	weapons := controls.NewWeapons()
	shields := controls.NewShields()
	scanner := controls.NewScanner()

	bridge.New(thruster, weapons, shields, scanner)

	thruster.SetThrust(90)
	weapons.ToggleArm()
	thruster.SetThrust(80)
	shields.ToggleShields()
	scanner.RequestScan()
}
