package model

type Item struct {
	Name string
	Tag  string
}

func (i Item) Clone() Item {
	return Item{Name: i.Name + "_clone", Tag: i.Tag}
}
