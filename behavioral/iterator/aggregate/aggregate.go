package aggregate

import (
	"iter"

	"github.com/kevinronu/dessign-patterns-go/behavioral/iterator/item"
)

// SocialNetwork returns sequences consumed with range. The compiler creates the
// yield callback for range and passes it to the sequence; the callback returns
// false when the loop stops early, such as with break.
// See https://go.dev/blog/range-functions#standard-push-iterators.
type SocialNetwork interface {
	FriendsFor(profileEmail string) iter.Seq[*item.Profile]
	CoworkersFor(profileEmail string) iter.Seq[*item.Profile]
}
