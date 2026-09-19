package builder

import (
	parta "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-a"
	partb "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-b"
	partc "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-c"
	"github.com/kevinronu/dessign-patterns-go/creational/builder/product"
)

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) BuildVariantOne() (product.Product, error) {
	partA := parta.PartA{Capacity: 3.0}

	d.builder.
		Reset().
		SetPartA(partA).
		SetPartB(partb.PartB{Profile: "default-profile"}).
		SetPartC(*partc.NewPartC(&partA))

	return d.builder.Build()
}

func (d *Director) BuildVariantTwo() (product.Product, error) {
	partA := parta.PartA{Capacity: 1.2}

	d.builder.
		Reset().
		SetPartA(partA).
		SetPartB(partb.PartB{Profile: "custom-profile"}).
		SetPartC(*partc.NewPartC(&partA))

	return d.builder.Build()
}

func (d *Director) BuildVariantThree() (product.Product, error) {
	d.builder.
		Reset().
		SetPartA(parta.PartA{Capacity: 2.5}).
		SetPartB(partb.PartB{Profile: "default-profile"})

	return d.builder.Build()
}
