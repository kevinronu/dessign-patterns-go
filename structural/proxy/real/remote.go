package real

import (
	"fmt"
	"time"
)

const latency = 300 * time.Millisecond

var videos = map[string]string{
	"cats": "🐱",
	"dogs": "🐶",
}

type Remote struct{}

func (Remote) Download(id string) ([]byte, error) {
	time.Sleep(latency)

	art, ok := videos[id]
	if !ok {
		return nil, fmt.Errorf("no video with id %q", id)
	}

	return []byte(art), nil
}
