package main

import (
	"fmt"
	"log"

	abstractfactory "github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/factory"
	"github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product"
	branda "github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product/brand-a"
	brandb "github.com/kevinronu/dessign-patterns-go/creational/abstract-factory/product/brand-b"
)

// printProductADetails prints a ProductA's brand and behavior.
func printProductADetails(p product.ProductA) {
	fmt.Printf("Brand: %s\n", p.GetBrand())
	fmt.Printf("Action: %s\n", p.DoA())
}

// printProductBDetails prints a ProductB's brand and behavior.
func printProductBDetails(p product.ProductB) {
	fmt.Printf("Brand: %s\n", p.GetBrand())
	fmt.Printf("Action: %s\n", p.DoB())
}

func main() {
	factoryA, err := abstractfactory.GetFactory(branda.BrandA)
	if err != nil {
		log.Fatalf("get factory A: %v", err)
	}

	factoryB, err := abstractfactory.GetFactory(brandb.BrandB)
	if err != nil {
		log.Fatalf("get factory B: %v", err)
	}

	productAFromA := factoryA.CreateProductA()
	productBFromA := factoryA.CreateProductB()

	productAFromB := factoryB.CreateProductA()
	productBFromB := factoryB.CreateProductB()

	printProductADetails(productAFromA)
	printProductBDetails(productBFromA)

	printProductADetails(productAFromB)
	printProductBDetails(productBFromB)

	fmt.Println(abstractfactory.OneTimeAction(factoryA))
	fmt.Println(abstractfactory.OneTimeAction(factoryB))
}
