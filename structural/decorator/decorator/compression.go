// Package decorator holds one wrapper for each behavior added to a DataSource.
//
// A wrapper keeps the value it wraps in a private wrappee field, and does not embed the
// interface. If DataSource gets a new method, this package stops compiling. With an embedded
// interface the new method would be forwarded unchanged, and nobody would notice.
package decorator

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"

	"github.com/kevinronu/dessign-patterns-go/structural/decorator/component"
)

func compress(data []byte, level int) ([]byte, error) {
	var buf bytes.Buffer

	writer, err := gzip.NewWriterLevel(&buf, level)
	if err != nil {
		return nil, fmt.Errorf("open gzip writer: %w", err)
	}

	if _, err := writer.Write(data); err != nil {
		return nil, fmt.Errorf("write gzip: %w", err)
	}

	// Close writes the gzip footer, so the buffer is not complete until it returns.
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("close gzip writer: %w", err)
	}

	return buf.Bytes(), nil
}

func decompress(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("open gzip reader: %w", err)
	}

	plain, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("read gzip: %w", err)
	}

	if err := reader.Close(); err != nil {
		return nil, fmt.Errorf("close gzip reader: %w", err)
	}

	return plain, nil
}

// Compression keeps the gzip work in the two functions above, so the methods below show only
// the pattern: change the data, then pass it down.
type Compression struct {
	wrappee component.DataSource
	level   int
}

// NewCompression needs no error because it only stores the level. gzip rejects a bad level on
// the first Write.
func NewCompression(wrappee component.DataSource, level int) Compression {
	return Compression{wrappee: wrappee, level: level}
}

func (c Compression) Write(data []byte) error {
	compressed, err := compress(data, c.level)
	if err != nil {
		return err
	}

	return c.wrappee.Write(compressed)
}

func (c Compression) Read() ([]byte, error) {
	data, err := c.wrappee.Read()
	if err != nil {
		// The layer below already said what went wrong. Adding context in every layer would give
		// a message like "a: b: c: EOF".
		return nil, err
	}

	return decompress(data)
}
