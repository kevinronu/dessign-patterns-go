package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kevinronu/dessign-patterns-go/creational/singleton/singleton"
)

func main() {
	ctx := context.Background()

	first, err := singleton.GetInstance(ctx)
	if err != nil {
		log.Fatalf("get instance: %v", err)
	}

	second, err := singleton.GetInstance(ctx)
	if err != nil {
		log.Fatalf("get instance: %v", err)
	}

	fmt.Println(first.Describe())
	fmt.Printf("same instance: %t\n", first == second)
}
