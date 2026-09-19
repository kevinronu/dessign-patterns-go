package productb

import (
	"fmt"
	"strings"

	parta "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-a"
	partb "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-b"
	partc "github.com/kevinronu/dessign-patterns-go/creational/builder/part/part-c"
	"github.com/kevinronu/dessign-patterns-go/creational/builder/product"
)

type ProductB struct {
	PartA *parta.PartA
	PartB *partb.PartB
	PartC *partc.PartC
}

var _ product.Product = ProductB{}

func (p ProductB) Describe() string {
	var b strings.Builder

	if p.PartA != nil {
		fmt.Fprintf(&b, "PartA: capacity %.1f\n", p.PartA.Capacity)
	} else {
		fmt.Fprintf(&b, "PartA: N/A\n")
	}

	if p.PartB != nil {
		fmt.Fprintf(&b, "PartB: %s\n", p.PartB.Profile)
	} else {
		fmt.Fprintf(&b, "PartB: N/A\n")
	}

	if p.PartC != nil {
		fmt.Fprintf(&b, "PartC: functional (%s)\n", p.PartC.ShowStatus())
	} else {
		fmt.Fprintf(&b, "PartC: N/A\n")
	}

	return b.String()
}
