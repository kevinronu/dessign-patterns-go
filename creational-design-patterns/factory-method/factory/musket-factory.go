package factory

import (
	"fmt"

	"github.com/kevinronu/design-patterns-go/creational-design-patterns/factory-method/gun"
)

type MusketFactory struct{}

func (m MusketFactory) createGun(owner string) gun.Gun {
	return &gun.Musket{
		Owner:   owner,
		GunType: gun.GunTypeMusket,
		Power:   2,
	}
}

func (mf MusketFactory) OneTimeAction(temporaryOwner string) {
	gun := mf.createGun(temporaryOwner)
	fmt.Printf("%s is temporarily using the Musket.\n", temporaryOwner)
	gun.Fire()
	gun.SpecialMove()
}
