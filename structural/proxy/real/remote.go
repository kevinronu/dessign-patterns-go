// Package real holds the service the proxy stands in for.
//
// The name shadows Go's real builtin inside any file that imports it, which only matters for code
// doing complex number arithmetic.
package real

import (
	"fmt"
	"time"
)

// latency is what makes a proxy worth having: every download costs this much.
const latency = 300 * time.Millisecond

// videos holds what a download brings. One emoji is enough to see that the bytes arrived.
var videos = map[string]string{
	"cats": "🐱",
	"dogs": "🐶",
}

// Remote is the real service. It keeps no state, so the only thing a proxy can save here is time.
type Remote struct{}

func (Remote) Download(id string) ([]byte, error) {
	time.Sleep(latency)

	art, ok := videos[id]
	if !ok {
		return nil, fmt.Errorf("no video with id %q", id)
	}

	return []byte(art), nil
}
