package adaptee

type LegacyService struct{}

func (LegacyService) Charge(dollars float64) bool {
	return dollars > 0
}
