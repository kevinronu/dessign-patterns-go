package gun

import "fmt"

type Musket struct {
	Owner   string
	Power   int
	GunType GunType
}

func (m Musket) Fire() {
	fmt.Printf("%s fired their %s and caused %d damage\n", m.Owner, m.GunType, m.Power)
}

func (m Musket) SpecialMove() {
	fmt.Printf("%s performs a precision shot with their %s, doubling the damage to %d\n", m.Owner, m.GunType, m.Power*2)
}

func NewMusket(owner string) Gun {
	return &Musket{
		Owner:   owner,
		GunType: GunTypeMusket,
		Power:   2,
	}
}
