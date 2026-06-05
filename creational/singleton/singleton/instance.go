package singleton

import "fmt"

// Instance is the single shared value this package exposes.
type Instance struct {
	name string
}

// Describe reports the instance identity.
func (i Instance) Describe() string {
	return fmt.Sprintf("instance %q", i.name)
}
