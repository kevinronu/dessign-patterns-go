package subsystem

type Raw struct {
	Bytes int64
}

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
