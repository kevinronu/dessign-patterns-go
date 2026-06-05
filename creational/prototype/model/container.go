package model

// Container is a prototype that owns a list of Entries.
type Container struct {
	Name    string
	Entries []Entry
}

// Clone returns a deep copy: every Entry is cloned and its back-reference is
// re-pointed to the new Container, so no entry still points at the original.
func (c *Container) Clone() *Container {
	clone := &Container{Name: c.Name + "_clone"}

	for _, entry := range c.Entries {
		cloned := entry.Clone()
		cloned.Container = clone
		clone.Entries = append(clone.Entries, cloned)
	}

	return clone
}
