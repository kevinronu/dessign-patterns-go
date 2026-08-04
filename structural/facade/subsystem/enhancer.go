package subsystem

// AudioEnhancer works on Raw and not on a MediaFile, because cleaning only makes sense while the
// audio is decoded, between the two codecs.
type AudioEnhancer struct{}

// Fix drops the quiet parts, so the result is smaller than a plain re-encode.
func (AudioEnhancer) Fix(raw Raw) Raw {
	return Raw{Bytes: raw.Bytes * 95 / 100}
}
