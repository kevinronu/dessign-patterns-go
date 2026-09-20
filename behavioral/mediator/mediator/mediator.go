package mediator

import "slices"

type Mediator interface {
	Register(component Component)
	Notify(sender Component, event Event)
}

type Component interface {
	Type() ComponentType
	SetMediator(mediator Mediator)
}

type ComponentType string

const (
	ThrusterType ComponentType = "thruster"
	WeaponsType  ComponentType = "weapons"
	ShieldsType  ComponentType = "shields"
	ScannerType  ComponentType = "scanner"
)

// Each event is a struct so its data travels typed, instead of a payload the
// mediator would have to assert.
type Event interface {
	Name() EventName
}

type EventName string

const (
	ThrustSetName      EventName = "SET_THRUST"
	WeaponsToggledName EventName = "TOGGLE_WEAPONS"
	ShieldsToggledName EventName = "TOGGLE_SHIELDS"
	ScanRequestedName  EventName = "RUN_SCAN"
)

type ThrustSet struct {
	Percent int
}

type WeaponsToggled struct {
	Armed bool
}

type ShieldsToggled struct {
	Up bool
}

type ScanRequested struct{}

func (ThrustSet) Name() EventName {
	return ThrustSetName
}

func (WeaponsToggled) Name() EventName {
	return WeaponsToggledName
}

func (ShieldsToggled) Name() EventName {
	return ShieldsToggledName
}

func (ScanRequested) Name() EventName {
	return ScanRequestedName
}

var allowedEvents = map[ComponentType][]EventName{
	ThrusterType: {ThrustSetName},
	WeaponsType:  {WeaponsToggledName},
	ShieldsType:  {ShieldsToggledName},
	ScannerType:  {ScanRequestedName},
}

// Any component can construct another component's event, so the mediator checks
// the pairing before reacting.
func Allows(component ComponentType, event EventName) bool {
	return slices.Contains(allowedEvents[component], event)
}
