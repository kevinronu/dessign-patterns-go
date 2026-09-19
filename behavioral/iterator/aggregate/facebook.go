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
		cachedProfiles := make([]*item.Profile, len(contactEmails))

		for position, contactEmail := range contactEmails {
			contact := cachedProfiles[position]
			if contact == nil {
				contact = f.requestProfile(contactEmail)
				cachedProfiles[position] = contact
			}

			keepGoing := yield(contact)
			if !keepGoing {
				return
			}
		}
	}
}
