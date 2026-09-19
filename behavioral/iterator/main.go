package main

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/behavioral/iterator/aggregate"
	"github.com/kevinronu/dessign-patterns-go/behavioral/iterator/client"
	"github.com/kevinronu/dessign-patterns-go/behavioral/iterator/item"
)

func main() {
	profiles := []item.Profile{
		item.New("anna@example.com", "Anna", "friends:max@example.com", "friends:sam@example.com", "coworkers:sam@example.com"),
		item.New("max@example.com", "Max"),
		item.New("sam@example.com", "Sam"),
	}

	facebookNetwork := aggregate.NewFacebook(profiles)
	facebookSpammer := client.New(facebookNetwork)
	linkedinSpammer := client.New(aggregate.NewLinkedIn(profiles))

	facebookSpammer.SendSpamToFriends("anna@example.com", "Please like this post")
	linkedinSpammer.SendSpamToCoworkers("anna@example.com", "Please read this update")

	reviewer := client.NewProfileReviewer(facebookNetwork.FriendsFor("anna@example.com"))
	defer reviewer.Stop()

	profile, ok := reviewer.Next()
	if !ok {
		return
	}

	fmt.Printf("reviewing %q\n", profile.Email)
	fmt.Println("user pauses the review")

	profile, ok = reviewer.Next()
	if !ok {
		return
	}

	fmt.Printf("reviewing %q\n", profile.Email)
}
