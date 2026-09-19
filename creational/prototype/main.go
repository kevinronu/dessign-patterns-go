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

	fmt.Printf("original: name=%q label=%q item=%q\n",
		container.Name, container.Entries[0].Label, container.Entries[0].Item.Name)
	fmt.Printf("clone:    name=%q label=%q item=%q\n",
		clonedContainer.Name, clonedContainer.Entries[0].Label, clonedContainer.Entries[0].Item.Name)

	fmt.Printf("back-reference: %q\n", clonedContainer.Entries[0].Container.Name)
}
