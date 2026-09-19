package client

import (
	"iter"

	"github.com/kevinronu/dessign-patterns-go/behavioral/iterator/item"
)

// ProfileReviewer lets a user review profiles one at a time.
type ProfileReviewer struct {
	Next func() (*item.Profile, bool)
	Stop func()
}

// NewProfileReviewer returns a reviewer that resumes from its last profile.
func NewProfileReviewer(profileSeq iter.Seq[*item.Profile]) *ProfileReviewer {
	next, stop := iter.Pull(profileSeq)

	return &ProfileReviewer{
		Next: next,
		Stop: stop,
	}
}
