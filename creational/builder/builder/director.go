package builder

import (
	parta "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-a"
	partb "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-b"
	partc "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-c"
	"github.com/kevinronu/dessign-patterns-go/creational/builder/product"
)

// Director drives the Builder it holds through a fixed recipe of steps.
type Director struct {
	builder Builder
}

// NewDirector binds a Director to a builder.
func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

// SetBuilder swaps the builder, so one recipe can drive any builder.
func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

// BuildVariantOne runs the first recipe and returns the product.
func (d *Director) BuildVariantOne() (product.Product, error) {
	partA := parta.PartA{Capacity: 3.0}

	d.builder.
		Reset().
		SetPartA(partA).
		SetPartB(partb.PartB{Profile: "default-profile"}).
		SetPartC(*partc.NewPartC(&partA))

	return d.builder.Build()
}

// BuildVariantTwo runs the second recipe and returns the product.
func (d *Director) BuildVariantTwo() (product.Product, error) {
	partA := parta.PartA{Capacity: 1.2}

	d.builder.
		Reset().
		SetPartA(partA).
		SetPartB(partb.PartB{Profile: "custom-profile"}).
		SetPartC(*partc.NewPartC(&partA))

	return d.builder.Build()
}

// BuildVariantThree runs the third recipe, which omits the optional PartC, and
// returns the product.
func (d *Director) BuildVariantThree() (product.Product, error) {
	d.builder.
		Reset().
		SetPartA(parta.PartA{Capacity: 2.5}).
		SetPartB(partb.PartB{Profile: "default-profile"})

	return d.builder.Build()
}
