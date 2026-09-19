package main

import (
	"fmt"
	"log"

	"github.com/kevinronu/dessign-patterns-go/structural/facade/facade"
	"github.com/kevinronu/dessign-patterns-go/structural/facade/subsystem"
)

func main() {
	source := subsystem.MediaFile{Name: "demo-video.ogg", Bytes: 614400}

	sourceCodec, err := subsystem.CodecFor(source.Format())
	if err != nil {
		log.Fatalf("source codec: %v", err)
	}

	targetCodec, err := subsystem.CodecFor("mp4")
	if err != nil {
		log.Fatalf("target codec: %v", err)
	}

	var (
		transcoder subsystem.Transcoder
		enhancer   subsystem.AudioEnhancer
	)

	raw := transcoder.Decode(source, sourceCodec)
	clean := enhancer.Fix(raw)
	byHand := transcoder.Encode(clean, targetCodec, source.BaseName())

	fmt.Println("by hand")
	fmt.Printf("  codec for %-3s    ratio %d\n", sourceCodec.Format, sourceCodec.Ratio)
	fmt.Printf("  codec for %-3s    ratio %d\n", targetCodec.Format, targetCodec.Ratio)
	fmt.Printf("  decode with %-3s  %7d B -> %7d B raw\n", sourceCodec.Format, source.Bytes, raw.Bytes)
	fmt.Printf("  fix              %7d B -> %7d B raw\n", raw.Bytes, clean.Bytes)
	fmt.Printf("  encode with %-3s  %7d B -> %7d B\n", targetCodec.Format, clean.Bytes, byHand.Bytes)
	fmt.Printf("  result           %s\n", byHand.Name)

	var converter facade.MediaConverter

	result, err := converter.Convert(source, "mp4")
	if err != nil {
		log.Fatalf("convert: %v", err)
	}

	fmt.Printf("\nthrough the facade\n  Convert          %s -> %s %d B\n", source.Name, result.Name, result.Bytes)

	fmt.Printf("\ncan convert to mp4\n")

	for _, name := range []string{"demo-video.ogg", "song.mp4", "clip.avi"} {
		if err := converter.Check(subsystem.MediaFile{Name: name}, "mp4"); err != nil {
			fmt.Printf("  %-15s no, %v\n", name, err)
			continue
		}

		fmt.Printf("  %-15s yes\n", name)
	}
}
