package main

import (
	"fmt"

	"github.com/kevinronu/design-patterns-go/creational-design-patterns/factory-method/factory"
	"github.com/kevinronu/design-patterns-go/creational-design-patterns/factory-method/gun"
)

func main() {
	ak47Factory, _ := factory.GetFactory(gun.GunTypeAK47)
	ak47Factory.OneTimeAction("Kevin")

	musketFactory, _ := factory.GetFactory(gun.GunTypeMusket)
	musketFactory.OneTimeAction("Noah")

	pistolFactory, err := factory.GetFactory("pistol")
	if err != nil {
		fmt.Println(err)

		return
	}

	pistolFactory.OneTimeAction("Leonardo")
}
