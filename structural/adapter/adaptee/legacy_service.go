// Package adaptee holds existing code whose API does not match the target.
package adaptee

// LegacyService is the adaptee: useful, but not ours to change, so we wrap it.
type LegacyService struct{}

func (LegacyService) Charge(dollars float64) bool {
	return dollars > 0
}
