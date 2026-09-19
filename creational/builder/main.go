package main

import (
	"fmt"
	"log"

	"github.com/kevinronu/dessign-patterns-go/creational/builder/builder"
)

func main() {
	builderA, err := builder.GetBuilder(builder.TypeA)
	if err != nil {
		log.Fatalf("get builder A: %v", err)
	}

	builderB, err := builder.GetBuilder(builder.TypeB)
	if err != nil {
		log.Fatalf("get builder B: %v", err)
	}

	director := builder.NewDirector(builderA)

	productA, err := director.BuildVariantOne()
	if err != nil {
		log.Fatalf("build product: %v", err)
	}

	fmt.Println(productA.Describe())

	director.SetBuilder(builderB)

	manual, err := director.BuildVariantTwo()
	if err != nil {
		log.Fatalf("build manual: %v", err)
	}

	fmt.Print(manual.Describe())
}
