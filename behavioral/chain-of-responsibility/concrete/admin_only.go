package concrete

import (
	"errors"
	"strings"

	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/handler"
	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/server"
)

type AdminOnly struct {
	handler.Successor

	server *server.Server
}

func NewAdminOnly(server *server.Server) *AdminOnly {
	return &AdminOnly{server: server}
}

func (a *AdminOnly) Handle(req handler.Request) error {
	if strings.HasPrefix(req.Path, "/admin") && !a.server.IsAdmin(req.Email) {
		return errors.New("admin only: not an admin")
	}

	return a.Successor.Handle(req)
}
