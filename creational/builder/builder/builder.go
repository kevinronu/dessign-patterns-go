package builder

import (
	"fmt"

	parta "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-a"
	partb "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-b"
	partc "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-c"
	"github.com/kevinronu/dessign-patterns-go/creational/builder/product"
)

// Builder assembles a product through chainable steps. Each setter returns the
// Builder so calls can be chained; Build completes assembly and returns the
// common Product, so concrete builders are interchangeable.
type Builder interface {
	Reset() Builder
	SetPartA(partA parta.PartA) Builder
	SetPartB(partB partb.PartB) Builder
	SetPartC(partC partc.PartC) Builder
	Build() (product.Product, error)
}

// BuilderType selects a concrete builder.
type BuilderType string

const (
	// TypeA builds a ProductA.
	TypeA BuilderType = "A"
	// TypeB builds a ProductB.
	TypeB BuilderType = "B"
)

// GetBuilder returns the concrete builder for the given type.
func GetBuilder(builderType BuilderType) (Builder, error) {
	switch builderType {
	case TypeA:
		return &BuilderA{}, nil
	case TypeB:
		return &BuilderB{}, nil
	default:
		return nil, fmt.Errorf("unsupported builder type: %q", builderType)
	}
}

// ValidateParts checks that every required part is set. PartC is optional.
func ValidateParts(partA *parta.PartA, partB *partb.PartB) error {
	if partA == nil {
		return fmt.Errorf("missing required part: PartA")
	}

	if partB == nil {
		return fmt.Errorf("missing required part: PartB")
	}

	return nil
}
