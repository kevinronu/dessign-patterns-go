package builder

import (
	parta "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-a"
	partb "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-b"
	partc "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-c"
	"github.com/kevinronu/dessign-patterns-go/creational/builder/product"
	producta "github.com/kevinronu/dessign-patterns-go/creational/builder/product/product-a"
)

// BuilderA is a Builder that produces a ProductA.
type BuilderA struct {
	partA *parta.PartA
	partB *partb.PartB
	partC *partc.PartC
}

// Reset clears accumulated parts for reuse.
func (b *BuilderA) Reset() Builder {
	*b = BuilderA{}

	return b
}

// SetPartA records PartA.
func (b *BuilderA) SetPartA(partA parta.PartA) Builder {
	b.partA = &partA

	return b
}

// SetPartB records PartB.
func (b *BuilderA) SetPartB(partB partb.PartB) Builder {
	b.partB = &partB

	return b
}

// SetPartC records PartC.
func (b *BuilderA) SetPartC(partC partc.PartC) Builder {
	b.partC = &partC

	return b
}

// Build returns the assembled ProductA, then resets the builder.
func (b *BuilderA) Build() (product.Product, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}

	result := producta.ProductA{
		PartA: b.partA,
		PartB: b.partB,
		PartC: b.partC,
	}

	b.Reset()

	return result, nil
}

// validate checks the required parts.
func (b *BuilderA) validate() error {
	return ValidateParts(b.partA, b.partB)
}
