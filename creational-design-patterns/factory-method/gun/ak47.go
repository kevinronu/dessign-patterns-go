package gun

import "fmt"

type AK47 struct {
	Owner   string
	Power   int
	GunType GunType
}

func (a AK47) Fire() {
	fmt.Printf("%s fired their %s and caused %d damage\n", a.Owner, a.GunType, a.Power)
}

func (a AK47) SpecialMove() {
	fmt.Printf("%s performs a burst fire with their %s, dealing %d damage rapidly\n", a.Owner, a.GunType, a.Power*3)
}
