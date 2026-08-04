package flyweight

import "sync"

// Factory keeps one flyweight per key. The lock is here because a cache is what several goroutines
// reach at once.
type Factory struct {
	mu    sync.RWMutex
	types map[string]*TreeType
}

func NewFactory() *Factory {
	return &Factory{types: make(map[string]*TreeType)}
}

// Get returns the flyweight for a key and builds it only the first time. It is safe for concurrent
// use.
func (f *Factory) Get(name, color, texture string) *TreeType {
	key := name + "|" + color + "|" + texture

	// Fast path: a hit needs only a read lock, and hits are what a cache mostly gets.
	f.mu.RLock()
	if shared, ok := f.types[key]; ok {
		f.mu.RUnlock()

		return shared
	}

	f.mu.RUnlock()

	// Slow path: Go cannot upgrade a read lock, so it has to be released first.
	f.mu.Lock()
	defer f.mu.Unlock()

	// Another goroutine may have created it while no lock was held.
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
