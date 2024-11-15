package factory

import (
	"fmt"

	"github.com/kevinronu/design-patterns-go/creational-design-patterns/factory-method/gun"
)

// Factory is the base interface for creating guns.
type Factory interface {
	createGun(owner string) gun.Gun
	OneTimeAction(temporaryOwner string)
}

func GetFactory(gunType gun.GunType) (Factory, error) {
	switch gunType {
	case gun.GunTypeAK47:
		return &AK47Factory{}, nil
	case gun.GunTypeMusket:
		return &MusketFactory{}, nil
	default:
		return nil, fmt.Errorf("unknown gun type: %s", gunType)
	}
}
