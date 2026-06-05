package producta

import (
	"fmt"
	"strings"

	parta "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-a"
	partb "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-b"
	partc "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-c"
	"github.com/kevinronu/dessign-patterns-go/creational/builder/product"
)

// ProductA is a product rendered as a single line.
type ProductA struct {
	PartA *parta.PartA
	PartB *partb.PartB
	PartC *partc.PartC
}

var _ product.Product = ProductA{}

// Describe renders the product as a single line.
func (p ProductA) Describe() string {
	var b strings.Builder

	b.WriteString("ProductA assembled with")

	if p.PartA != nil {
		fmt.Fprintf(&b, " PartA(capacity=%.1f)", p.PartA.Capacity)
	}

	if p.PartB != nil {
		fmt.Fprintf(&b, " PartB(%s)", p.PartB.Profile)
	}

	if p.PartC != nil {
		fmt.Fprintf(&b, " PartC(%s)", p.PartC.ShowStatus())
	}

	return b.String()
}
