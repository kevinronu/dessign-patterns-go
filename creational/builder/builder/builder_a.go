package builder

import (
	parta "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-a"
	partb "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-b"
	partc "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-c"
	"github.com/kevinronu/dessign-patterns-go/creational/builder/product"
	producta "github.com/kevinronu/dessign-patterns-go/creational/builder/product/product-a"
)

type BuilderA struct {
	partA *parta.PartA
	partB *partb.PartB
	partC *partc.PartC
}

func (b *BuilderA) Reset() Builder {
	*b = BuilderA{}

	return b
}

func (b *BuilderA) SetPartA(partA parta.PartA) Builder {
	b.partA = &partA

	return b
}

func (b *BuilderA) SetPartB(partB partb.PartB) Builder {
	b.partB = &partB

	return b
}

func (b *BuilderA) SetPartC(partC partc.PartC) Builder {
	b.partC = &partC

	return b
}

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

func (b *BuilderA) validate() error {
	return ValidateParts(b.partA, b.partB)
}
