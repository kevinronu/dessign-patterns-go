package singleton

import (
	"context"
	"fmt"
	"sync"
)

// instanceSingleton holds the lazily-created Instance behind a read/write lock.
type instanceSingleton struct {
	instance *Instance
	mu       sync.RWMutex
}

// shared is the package-level singleton holder.
var shared instanceSingleton

// GetInstance returns the shared Instance, creating it once on first use.
// It is safe for concurrent use.
func GetInstance(ctx context.Context) (*Instance, error) {
	// Fast path: return the instance if it already exists.
	shared.mu.RLock()
	if shared.instance != nil {
		shared.mu.RUnlock()

		return shared.instance, nil
	}

	shared.mu.RUnlock()

	// Slow path: take the write lock to create it.
	shared.mu.Lock()
	defer shared.mu.Unlock()

	// Re-check: another goroutine may have created it between the locks.
	if shared.instance != nil {
		return shared.instance, nil
	}

	created, err := buildInstance(ctx)
	if err != nil {
		return nil, fmt.Errorf("building instance: %w", err)
	}

	shared.instance = created

	return shared.instance, nil
}

// buildInstance performs the one-time setup that may fail.
func buildInstance(ctx context.Context) (*Instance, error) {
	// A real build might read a secret or open a connection here.
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return &Instance{name: "shared"}, nil
}
