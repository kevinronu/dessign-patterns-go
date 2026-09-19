package client

import (
	"fmt"
	"iter"

	"github.com/kevinronu/dessign-patterns-go/behavioral/iterator/aggregate"
	"github.com/kevinronu/dessign-patterns-go/behavioral/iterator/item"
)

type SocialSpammer struct {
	network aggregate.SocialNetwork
}

func New(network aggregate.SocialNetwork) SocialSpammer {
	return SocialSpammer{network: network}
}

func (s SocialSpammer) SendSpamToFriends(email, message string) {
	s.send(s.network.FriendsFor(email), message)
}

func (s SocialSpammer) SendSpamToCoworkers(email, message string) {
	s.send(s.network.CoworkersFor(email), message)
}

func (SocialSpammer) send(profileSeq iter.Seq[*item.Profile], message string) {
	for profile := range profileSeq {
		if profile != nil {
			fmt.Printf("sent to %q: %s\n", profile.Email, message)
		}
	}
}
