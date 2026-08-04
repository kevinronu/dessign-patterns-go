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

	// Two documents and two formats give four outputs from four small types. Without the split,
	// every pair would need a type of its own.
	for _, exporter := range []implementation.Exporter{implementation.HTML{}, implementation.CSV{}} {
		// The Exporter is a field, so an implementation can be swapped after the document is built.
		invoice.Exporter = exporter
		report.Exporter = exporter

		fmt.Printf("--- invoice as %T ---\n%s\n", exporter, invoice.Export())
		fmt.Printf("--- report as %T ---\n%s\n", exporter, report.Export())
	}
}
