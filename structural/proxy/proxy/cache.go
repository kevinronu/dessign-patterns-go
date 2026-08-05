// Package proxy holds the stand-in the client talks to instead of the real service.
package proxy

import (
	"bytes"
	"sync"
	"time"

	"github.com/kevinronu/dessign-patterns-go/structural/proxy/service"
)

// Cache is the proxy: the same contract as the service, but it only asks the service for an id it
// has not seen. The service needs no changes and never learns the proxy exists.
type Cache struct {
	mu     sync.RWMutex
	origin service.VideoLibrary
	stored map[string][]byte
	stop   func()
}

// NewCache starts the cleanup goroutine, so every Cache has to be stopped. That is the price of a
// background task in Go: nothing collects a running goroutine for you.
//
// ttl is a cleanup interval, not a lifetime per entry: every tick drops everything, no matter when
// it got there.
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

// clean waits on two channels at once: the tick that empties the cache, and the order to finish.
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

// Stop ends the cleanup goroutine. It goes through sync.OnceFunc because closing a channel twice
// panics.
func (c *Cache) Stop() {
	c.stop()
}

func (c *Cache) Download(id string) ([]byte, error) {
	// Fast path: a stored answer needs only a read lock.
	c.mu.RLock()
	if art, ok := c.stored[id]; ok {
		defer c.mu.RUnlock()

		// The bytes go out as a copy, so a caller cannot change what the next one gets.
		return bytes.Clone(art), nil
	}

	c.mu.RUnlock()

	// Slow path: Go cannot upgrade a read lock, so it has to be released first.
	c.mu.Lock()
	defer c.mu.Unlock()

	// Another goroutine may have downloaded it while no lock was held.
	if art, ok := c.stored[id]; ok {
		return bytes.Clone(art), nil
	}

	art, err := c.origin.Download(id)
	if err != nil {
		// A failure is not stored, so the next call asks again.
		return nil, err
	}

	c.stored[id] = art

	return bytes.Clone(art), nil
}
