package proxy

import (
	"bytes"
	"sync"
	"time"

	"github.com/kevinronu/dessign-patterns-go/structural/proxy/service"
)

type Cache struct {
	mu     sync.RWMutex
	origin service.VideoLibrary
	stored map[string][]byte
	stop   func()
}

// NewCache clears its full cache every ttl and starts a goroutine that Stop must end.
func NewCache(origin service.VideoLibrary, ttl time.Duration) *Cache {
	done := make(chan struct{})

	cache := &Cache{
		origin: origin,
		stored: make(map[string][]byte),
		stop:   sync.OnceFunc(func() { close(done) }),
	}

	go cache.clean(ttl, done)

	return cache
}

func (c *Cache) clean(ttl time.Duration, done <-chan struct{}) {
	ticker := time.NewTicker(ttl)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			clear(c.stored)
			c.mu.Unlock()
		case <-done:
			return
		}
	}
}

func (c *Cache) Stop() {
	c.stop()
}

func (c *Cache) Download(id string) ([]byte, error) {
	c.mu.RLock()
	if art, ok := c.stored[id]; ok {
		defer c.mu.RUnlock()

		return bytes.Clone(art), nil
	}

	c.mu.RUnlock()

	c.mu.Lock()
	defer c.mu.Unlock()

	if art, ok := c.stored[id]; ok {
		return bytes.Clone(art), nil
	}

	art, err := c.origin.Download(id)
	if err != nil {
		return nil, err
	}

	c.stored[id] = art

	return bytes.Clone(art), nil
}
