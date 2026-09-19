package flyweight

import "sync"

type Factory struct {
	mu    sync.RWMutex
	types map[string]*TreeType
}

func NewFactory() *Factory {
	return &Factory{types: make(map[string]*TreeType)}
}

func (f *Factory) Get(name, color, texture string) *TreeType {
	key := name + "|" + color + "|" + texture

	f.mu.RLock()
	if shared, ok := f.types[key]; ok {
		f.mu.RUnlock()

		return shared
	}

	f.mu.RUnlock()

	f.mu.Lock()
	defer f.mu.Unlock()

	if shared, ok := f.types[key]; ok {
		return shared
	}

	shared := &TreeType{name: name, color: color, texture: texture}
	f.types[key] = shared

	return shared
}

func (f *Factory) Size() int {
	f.mu.RLock()
	defer f.mu.RUnlock()

	return len(f.types)
}
