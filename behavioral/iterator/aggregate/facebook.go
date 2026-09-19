package aggregate

import (
	"iter"

	"github.com/kevinronu/dessign-patterns-go/behavioral/iterator/item"
)

type Facebook struct {
	profiles []item.Profile
}

func NewFacebook(profiles []item.Profile) *Facebook {
	return &Facebook{profiles: profiles}
}

func (f *Facebook) requestProfile(email string) *item.Profile {
	// Use the index because range values are copies of slice elements.
	for index := range f.profiles {
		if f.profiles[index].Email == email {
			return &f.profiles[index]
		}
	}

	return nil
}

func (f *Facebook) requestProfileContactEmails(email string, contactType item.ContactType) []string {
	p := f.requestProfile(email)
	if p == nil {
		return nil
	}

	return p.Contacts[contactType]
}

func (f *Facebook) FriendsFor(profileEmail string) iter.Seq[*item.Profile] {
	return f.profileSeqFor(profileEmail, item.Friend)
}

func (f *Facebook) CoworkersFor(profileEmail string) iter.Seq[*item.Profile] {
	return f.profileSeqFor(profileEmail, item.Coworker)
}

func (f *Facebook) profileSeqFor(email string, contactType item.ContactType) iter.Seq[*item.Profile] {
	return func(yield func(*item.Profile) bool) {
		contactEmails := f.requestProfileContactEmails(email, contactType)

		for _, contactEmail := range contactEmails {
			contact := f.requestProfile(contactEmail)

			keepGoing := yield(contact)
			if !keepGoing {
				return
			}
		}
	}
}
