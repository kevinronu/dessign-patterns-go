package main

import (
	"fmt"
	"log"

	"github.com/kevinronu/dessign-patterns-go/structural/composite/composite"
	"github.com/kevinronu/dessign-patterns-go/structural/composite/leaf"
)

func main() {
	docs := &composite.Folder{Title: "docs"}
	docs.Add(
		leaf.File{Title: "report.pdf", Bytes: 2048},
		leaf.File{Title: "notes.md", Bytes: 320},
		leaf.Symlink{Title: "latest", Target: "report.pdf"},
	)

	img := &composite.Folder{Title: "img"}
	img.Add(leaf.File{Title: "logo.png", Bytes: 512})

	root := &composite.Folder{Title: "project"}
	root.Add(docs, img)

	fmt.Printf("%s totals %d B\n\n", root.Name(), root.Size())

	paths := []string{
		"project/docs/report.pdf",
		"project/docs/latest",
		"project/docs",
		"project/docs/missing.txt",
	}

	for _, path := range paths {
		found, ok := root.Find(path)
		if !ok {
			fmt.Printf("%-26s not found\n", path)
			continue
		}

		fmt.Printf("%-26s %-12s %6d B\n", path, found.Name(), found.Size())
	}

	if err := docs.Remove("report.pdf"); err != nil {
		log.Fatalf("remove report.pdf: %v", err)
	}

	fmt.Printf("\n%s totals %d B after removing report.pdf\n", root.Name(), root.Size())
}
