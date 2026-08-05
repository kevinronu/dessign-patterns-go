// Package concrete holds the steps of the chain. Each one refuses for its own reason and knows
// nothing about the others.
package concrete

import (
	"fmt"
	"time"

	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/handler"
)

// RateLimit spaces requests out with GCRA: it keeps no counter, only the time when the next request
// is due. GCRA calls that time the theoretical arrival time.
//
// A window counter is easier to read, but a window has to be reset, and five requests just before
// the reset plus five just after make ten in a moment. A schedule is never reset.
//
// One request at a time, like the demo. A chain shared by goroutines would need a lock here.
type RateLimit struct {
	handler.Successor

	interval time.Duration
	burst    time.Duration // how early a caller may arrive, so requests can come together
	due      time.Time
}

func NewRateLimit(limit int, window time.Duration) *RateLimit {
	return &RateLimit{interval: window / time.Duration(limit), burst: window}
}

func (r *RateLimit) Handle(req handler.Request) error {
	now := time.Now()

	// An idle limiter falls behind. Catching up to now caps what it earned at one burst, and the
	// zero time lands here too, so a fresh limiter needs no setup.
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
