package item

import "strings"

type ContactType string

const (
	Friend   ContactType = "friends"
	Coworker ContactType = "coworkers"
)

type Profile struct {
	Email    string
	Name     string
	Contacts map[ContactType][]string
}

func New(email, name string, contacts ...string) Profile {
	p := Profile{
		Email:    email,
		Name:     name,
		Contacts: make(map[ContactType][]string),
	}

	for _, contact := range contacts {
		kind, contactEmail, found := strings.Cut(contact, ":")
		if !found {
			kind, contactEmail = string(Friend), contact
		}

		contactType := ContactType(kind)
		p.Contacts[contactType] = append(p.Contacts[contactType], contactEmail)
	}

	return p
}
