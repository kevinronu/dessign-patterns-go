package singleton

import (
	"context"
	"fmt"
	"sync"
)

type instanceSingleton struct {
	instance *Instance
	mu       sync.RWMutex
}

var shared instanceSingleton

// GetInstance returns the process-wide instance, creating it with ctx on the first successful call.
// It is safe for concurrent use. A failed initialization leaves no instance and later calls retry.
func GetInstance(ctx context.Context) (*Instance, error) {
	shared.mu.RLock()
	if shared.instance != nil {
		shared.mu.RUnlock()

		return shared.instance, nil
	}

	shared.mu.RUnlock()

	shared.mu.Lock()
	defer shared.mu.Unlock()

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

func buildInstance(ctx context.Context) (*Instance, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return &Instance{name: "shared"}, nil
}
