package model

// Entry is a composite prototype: it owns an Item and an optional
// back-reference to its Container.
type Entry struct {
	Label     string     // copied as-is
	Item      Item       // nested prototype, deep-cloned
	Data      string     // copied as-is
	Container *Container // back-reference, re-linked by Container.Clone
}

// Clone returns a deep copy with no shared references. The Container
// back-reference is left nil; Container.Clone re-links it.
func (e Entry) Clone() Entry {
	return Entry{
		Label:     e.Label + "_clone",
		Item:      e.Item.Clone(),
		Data:      e.Data,
		Container: nil,
	}
}
