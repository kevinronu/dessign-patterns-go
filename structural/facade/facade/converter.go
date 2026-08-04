// Package facade holds the only type the client needs to know.
package facade

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/structural/facade/subsystem"
)

// MediaConverter owns the subsystem parts, so the client never builds them. Both are zero value
// structs, which is why a MediaConverter needs no constructor.
type MediaConverter struct {
	transcoder subsystem.Transcoder
	enhancer   subsystem.AudioEnhancer
}

// codecs pairs the two ends of a conversion, so a failed lookup returns one zero value instead of
// two.
type codecs struct {
	source subsystem.Codec
	target subsystem.Codec
}

// codecPair is the step both methods below start with. It lives here so neither of them repeats it.
func codecPair(file subsystem.MediaFile, format string) (codecs, error) {
	source, err := subsystem.CodecFor(file.Format())
	if err != nil {
		return codecs{}, fmt.Errorf("source codec: %w", err)
	}

	target, err := subsystem.CodecFor(format)
	if err != nil {
		return codecs{}, fmt.Errorf("target codec: %w", err)
	}

	return codecs{source: source, target: target}, nil
}

// Check stops after the codecs, so a client can reject a file before anything is decoded. The error
// says which end is unsupported.
func (MediaConverter) Check(file subsystem.MediaFile, format string) error {
	_, err := codecPair(file, format)

	return err
}

// Convert replaces the five calls the subsystem needs in the right order. A facade picks the parts
// and the order, it does not add work of its own.
func (c MediaConverter) Convert(file subsystem.MediaFile, format string) (subsystem.MediaFile, error) {
	pair, err := codecPair(file, format)
	if err != nil {
		return subsystem.MediaFile{}, err
	}

	raw := c.transcoder.Decode(file, pair.source)

	return c.transcoder.Encode(c.enhancer.Fix(raw), pair.target, file.BaseName()), nil
}
