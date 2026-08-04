// Package subsystem holds the parts that do the real work. Each one knows its own job and nothing
// about the order they run in.
package subsystem

import "fmt"

// Codec is one part of the subsystem. Ratio is how many raw bytes fit into one stored byte.
type Codec struct {
	Format string
	Ratio  int
}

// codecs is a read-only registry, so a new format never touches the facade.
var codecs = map[string]Codec{
	"ogg": {Format: "ogg", Ratio: 12},
	"mp4": {Format: "mp4", Ratio: 20},
}

// CodecFor fails on an unknown format, which is the one error the facade has to pass on.
func CodecFor(format string) (Codec, error) {
	codec, ok := codecs[format]
	if !ok {
		return Codec{}, fmt.Errorf("no codec for format %q", format)
	}

	return codec, nil
}
