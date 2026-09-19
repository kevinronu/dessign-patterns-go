package facade

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/structural/facade/subsystem"
)

type MediaConverter struct {
	transcoder subsystem.Transcoder
	enhancer   subsystem.AudioEnhancer
}

type codecs struct {
	source subsystem.Codec
	target subsystem.Codec
}

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

func (MediaConverter) Check(file subsystem.MediaFile, format string) error {
	_, err := codecPair(file, format)

	return err
}

func (c MediaConverter) Convert(file subsystem.MediaFile, format string) (subsystem.MediaFile, error) {
	pair, err := codecPair(file, format)
	if err != nil {
		return subsystem.MediaFile{}, err
	}

	raw := c.transcoder.Decode(file, pair.source)

	return c.transcoder.Encode(c.enhancer.Fix(raw), pair.target, file.BaseName()), nil
}
