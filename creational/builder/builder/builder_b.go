package builder

import (
	parta "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-a"
	partb "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-b"
	partc "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-c"
	"github.com/kevinronu/dessign-patterns-go/creational/builder/product"
	productb "github.com/kevinronu/dessign-patterns-go/creational/builder/product/product-b"
)

// BuilderB is a Builder that produces a ProductB.
type BuilderB struct {
	partA *parta.PartA
	partB *partb.PartB
	partC *partc.PartC
}

// Reset clears accumulated parts for reuse.
func (b *BuilderB) Reset() Builder {
	*b = BuilderB{}

	return b
}

// SetPartA records PartA.
func (b *BuilderB) SetPartA(partA parta.PartA) Builder {
	b.partA = &partA

	return b
}

// SetPartB records PartB.
func (b *BuilderB) SetPartB(partB partb.PartB) Builder {
	b.partB = &partB

	return b
}

// SetPartC records PartC.
func (b *BuilderB) SetPartC(partC partc.PartC) Builder {
	b.partC = &partC

	return b
}

// Build returns the assembled ProductB, then resets the builder.
func (b *BuilderB) Build() (product.Product, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}

	result := productb.ProductB{
		PartA: b.partA,
		PartB: b.partB,
		PartC: b.partC,
	}

	b.Reset()

	return result, nil
}

// validate checks the required parts.
func (b BuilderB) validate() error {
	return ValidateParts(b.partA, b.partB)
}
