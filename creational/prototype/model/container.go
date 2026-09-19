package model

type Container struct {
	Name    string
	Entries []Entry
}

// Clone returns a deep copy whose entries refer to the clone, not the original container.
func (c *Container) Clone() *Container {
	clone := &Container{Name: c.Name + "_clone"}

	for _, entry := range c.Entries {
		cloned := entry.Clone()
		cloned.Container = clone
		clone.Entries = append(clone.Entries, cloned)
	}

	return clone
}
