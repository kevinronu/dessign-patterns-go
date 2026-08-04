package subsystem

// Raw is decoded audio. It is a separate type so the two steps cannot be called in the wrong
// order: nothing else in the subsystem accepts or returns it.
type Raw struct {
	Bytes int64
}

// Transcoder holds no state, so the facade can keep it as a zero value.
type Transcoder struct{}

func (Transcoder) Decode(file MediaFile, codec Codec) Raw {
	return Raw{Bytes: file.Bytes * int64(codec.Ratio)}
}

func (Transcoder) Encode(raw Raw, codec Codec, baseName string) MediaFile {
	return MediaFile{
		Name:  baseName + "." + codec.Format,
		Bytes: raw.Bytes / int64(codec.Ratio),
	}
}
