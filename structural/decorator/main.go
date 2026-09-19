package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"strings"

	"github.com/kevinronu/dessign-patterns-go/structural/decorator/component"
	"github.com/kevinronu/dessign-patterns-go/structural/decorator/concrete"
	"github.com/kevinronu/dessign-patterns-go/structural/decorator/decorator"
)

func main() {
	original := []byte("name,salary\n" + strings.Repeat("John Smith,100000\nSteven Jobs,912000\n", 20))

	const level = gzip.BestCompression

	key := []byte("0123456789abcdef0123456789abcdef")

	store := &concrete.Store{}

	encrypted, err := decorator.NewEncryption(store, key)
	if err != nil {
		log.Fatalf("new encryption: %v", err)
	}

	swapped, err := decorator.NewEncryption(decorator.NewCompression(store, level), key)
	if err != nil {
		log.Fatalf("new encryption: %v", err)
	}

	stacks := []struct {
		label  string
		source component.DataSource
	}{
		{"none", store},
		{"compression", decorator.NewCompression(store, level)},
		{"compression + encryption", decorator.NewCompression(encrypted, level)},
		{"encryption + compression", swapped},
	}

	fmt.Printf("%-26s %8s   %s\n", "layers", "in store", "same after read")

	for _, s := range stacks {
		if err := s.source.Write(original); err != nil {
			log.Fatalf("%s: write: %v", s.label, err)
		}

		restored, err := s.source.Read()
		if err != nil {
			log.Fatalf("%s: read: %v", s.label, err)
		}

		fmt.Printf("%-26s %6d B   %t\n", s.label, len(store.Data), bytes.Equal(original, restored))
	}
}
