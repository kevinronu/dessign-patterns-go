package subsystem

import "fmt"

type Codec struct {
	Format string
	Ratio  int
}

var codecs = map[string]Codec{
	"ogg": {Format: "ogg", Ratio: 12},
	"mp4": {Format: "mp4", Ratio: 20},
}

func CodecFor(format string) (Codec, error) {
	codec, ok := codecs[format]
	if !ok {
		return Codec{}, fmt.Errorf("no codec for format %q", format)
	}

	return codec, nil
}
