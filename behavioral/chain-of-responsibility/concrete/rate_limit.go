package concrete

import (
	"fmt"
	"time"

	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/handler"
)

type RateLimit struct {
	handler.Successor

	interval time.Duration
	burst    time.Duration
	due      time.Time
}

// NewRateLimit allows limit requests per window on average, with an initial burst of one window.
func NewRateLimit(limit int, window time.Duration) *RateLimit {
	return &RateLimit{interval: window / time.Duration(limit), burst: window}
}

func (r *RateLimit) Handle(req handler.Request) error {
	now := time.Now()

	due := r.due
	if due.Before(now) {
		due = now
	}

	if allowAt := due.Add(r.interval - r.burst); now.Before(allowAt) {
		return fmt.Errorf("rate limit: too many requests, wait %v",
			allowAt.Sub(now).Round(time.Millisecond))
	}

	r.due = due.Add(r.interval)

	return r.Successor.Handle(req)
}
