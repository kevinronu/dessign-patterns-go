package main

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/creational/prototype/model"
)

func main() {
	container := &model.Container{Name: "base"}
	container.Entries = []model.Entry{
		{Label: "first", Item: model.Item{Name: "alpha", Tag: "x"}, Data: "d1"},
		{Label: "second", Item: model.Item{Name: "beta", Tag: "y"}},
	}

	clonedContainer := container.Clone()

	// The clone carries a "_clone" suffix on every field: it is a separate deep copy.
	fmt.Printf("original: name=%q label=%q item=%q\n",
		container.Name, container.Entries[0].Label, container.Entries[0].Item.Name)
	fmt.Printf("clone:    name=%q label=%q item=%q\n",
		clonedContainer.Name, clonedContainer.Entries[0].Label, clonedContainer.Entries[0].Item.Name)

	// The clone's back-reference points to the clone, not the original.
	fmt.Printf("back-reference: %q\n", clonedContainer.Entries[0].Container.Name)
}
