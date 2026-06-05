package main

import (
	"fmt"
	"log"

	factorymethod "github.com/kevinronu/dessign-patterns-go/creational/factory-method/factory"
	"github.com/kevinronu/dessign-patterns-go/creational/factory-method/product"
)

// printDetails prints the type and behavior of a product.
func printDetails(p product.Product) {
	fmt.Printf("Type: %s\n", p.GetType())
	fmt.Printf("Action: %s\n", p.DoSomething())
}

func main() {
	factoryA, err := factorymethod.GetFactory(product.TypeA)
	if err != nil {
		log.Fatalf("get factory A: %v", err)
	}

	factoryB, err := factorymethod.GetFactory(product.TypeB)
	if err != nil {
		log.Fatalf("get factory B: %v", err)
	}

	productA := factoryA.CreateProduct("Alice")
	productB := factoryB.CreateProduct("Bob")

	printDetails(productA)
	printDetails(productB)

	fmt.Println(factorymethod.OneTimeAction(factoryA, "Carol"))
	fmt.Println(factorymethod.OneTimeAction(factoryB, "Dave"))
}
