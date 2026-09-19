package main

import (
	"github.com/kevinronu/dessign-patterns-go/behavioral/iterator/aggregate"
	"github.com/kevinronu/dessign-patterns-go/behavioral/iterator/client"
	"github.com/kevinronu/dessign-patterns-go/behavioral/iterator/item"
)

func main() {
	profiles := []item.Profile{
		item.New("anna@example.com", "Anna", "friends:max@example.com", "coworkers:sam@example.com"),
		item.New("max@example.com", "Max"),
		item.New("sam@example.com", "Sam"),
	}

	facebook := client.New(aggregate.NewFacebook(profiles))
	linkedin := client.New(aggregate.NewLinkedIn(profiles))

	facebook.SendSpamToFriends("anna@example.com", "Please like this post")
	linkedin.SendSpamToCoworkers("anna@example.com", "Please read this update")
}
