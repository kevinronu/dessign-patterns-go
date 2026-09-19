package model

type Entry struct {
	Label     string
	Item      Item
	Data      string
	Container *Container
}

// Clone returns a deep copy with no container. Container.Clone restores that back-reference.
func (e Entry) Clone() Entry {
	return Entry{
		Label:     e.Label + "_clone",
		Item:      e.Item.Clone(),
		Data:      e.Data,
		Container: nil,
	}
}
