package main

import (
	"fmt"
	"log"

	"github.com/kevinronu/dessign-patterns-go/structural/adapter/adaptee"
	"github.com/kevinronu/dessign-patterns-go/structural/adapter/adapter"
	"github.com/kevinronu/dessign-patterns-go/structural/adapter/client"
	"github.com/kevinronu/dessign-patterns-go/structural/adapter/native"
)

func main() {
	// A backend that already fits the target is used directly.
	nativePayment := client.PaymentService{Processor: native.Service{}}

	receipt, err := nativePayment.Buy(1999)
	if err != nil {
		log.Fatalf("native payment: %v", err)
	}

	fmt.Println("no adapter: ", receipt)

	// A backend whose API does not match the target needs an adapter.
	legacyPayment := client.PaymentService{Processor: adapter.LegacyAdapter{Adaptee: adaptee.LegacyService{}}}

	receipt, err = legacyPayment.Buy(2500)
	if err != nil {
		log.Fatalf("legacy payment: %v", err)
	}

	fmt.Println("via adapter:", receipt)
}
