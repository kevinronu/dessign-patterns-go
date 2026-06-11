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

	// Render the same documents with each exporter — swap the implementation, reuse the data.
	for _, exporter := range []implementation.Exporter{implementation.HTML{}, implementation.CSV{}} {
		invoice.Exporter = exporter
		report.Exporter = exporter

		fmt.Print(invoice.Export())
		fmt.Println()
		fmt.Print(report.Export())
		fmt.Println()
	}
}
