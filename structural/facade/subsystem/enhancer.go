package subsystem

type AudioEnhancer struct{}

func (AudioEnhancer) Fix(raw Raw) Raw {
	return Raw{Bytes: raw.Bytes * 95 / 100}
}
