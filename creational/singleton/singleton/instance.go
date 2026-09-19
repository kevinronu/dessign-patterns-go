package singleton

import "fmt"

type Instance struct {
	name string
}

func (i Instance) Describe() string {
	return fmt.Sprintf("instance %q", i.name)
}
