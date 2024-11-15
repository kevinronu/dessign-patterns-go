package factory

import (
	"fmt"

	"github.com/kevinronu/design-patterns-go/creational-design-patterns/factory-method/gun"
)

type AK47Factory struct{}

func (a AK47Factory) createGun(owner string) gun.Gun {
	return &gun.AK47{
		Owner:   owner,
		GunType: gun.GunTypeAK47,
		Power:   4,
	}
}

func (af AK47Factory) OneTimeAction(temporaryOwner string) {
	gun := af.createGun(temporaryOwner)
	fmt.Printf("%s is temporarily using the AK47.\n", temporaryOwner)
	gun.Fire()
	gun.SpecialMove()
}
