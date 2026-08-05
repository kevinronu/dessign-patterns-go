package concrete

import (
	"errors"

	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/handler"
	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/server"
)

// Auth proves who is calling. Every request carries the credentials, the way basic auth does.
type Auth struct {
	handler.Successor

	server *server.Server
}

func NewAuth(server *server.Server) *Auth {
	return &Auth{server: server}
}

// Every error names this step, so the output shows how far the request got.
func (a *Auth) Handle(req handler.Request) error {
	if !a.server.HasEmail(req.Email) {
		return errors.New("auth: no such account")
	}

	if !a.server.IsValidPassword(req.Email, req.Password) {
		return errors.New("auth: wrong password")
	}

	return a.Successor.Handle(req)
}
