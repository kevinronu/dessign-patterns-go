package aggregate

import (
	"iter"

	"github.com/kevinronu/dessign-patterns-go/behavioral/iterator/item"
)

type LinkedIn struct {
	contacts []item.Profile
}

func NewLinkedIn(contacts []item.Profile) *LinkedIn {
	return &LinkedIn{contacts: contacts}
}

func (l *LinkedIn) requestProfile(email string) *item.Profile {
	// Use the index because range values are copies of slice elements.
	for index := range l.contacts {
		if l.contacts[index].Email == email {
			return &l.contacts[index]
		}
	}

	return nil
}

func (l *LinkedIn) requestProfileContactEmails(email string, contactType item.ContactType) []string {
	p := l.requestProfile(email)
	if p == nil {
		return nil
	}

	return p.Contacts[contactType]
}

func (l *LinkedIn) FriendsFor(profileEmail string) iter.Seq[*item.Profile] {
	return l.profileSeqFor(profileEmail, item.Friend)
}

func (l *LinkedIn) CoworkersFor(profileEmail string) iter.Seq[*item.Profile] {
	return l.profileSeqFor(profileEmail, item.Coworker)
}

func (l *LinkedIn) profileSeqFor(email string, contactType item.ContactType) iter.Seq[*item.Profile] {
	return func(yield func(*item.Profile) bool) {
		contactEmails := l.requestProfileContactEmails(email, contactType)

		for _, contactEmail := range contactEmails {
			contact := l.requestProfile(contactEmail)

			keepGoing := yield(contact)
			if !keepGoing {
				return
			}
		}
	}
}
