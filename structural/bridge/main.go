package main

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/structural/bridge/abstraction"
	"github.com/kevinronu/dessign-patterns-go/structural/bridge/implementation"
)

func main() {
	invoice := abstraction.Invoice{
		Number:   "001",
		Customer: "Ana",
		Date:     "2026-06-10",
		Total:    "$50",
	}

	report := abstraction.Report{
		Title:   "Q1 Report",
		Author:  "Bob",
		Date:    "2026-06-10",
		Summary: "all good",
	}

	for _, exporter := range []implementation.Exporter{implementation.HTML{}, implementation.CSV{}} {
		invoice.Exporter = exporter
		report.Exporter = exporter

		fmt.Printf("--- invoice as %T ---\n%s\n", exporter, invoice.Export())
		fmt.Printf("--- report as %T ---\n%s\n", exporter, report.Export())
	}
}
