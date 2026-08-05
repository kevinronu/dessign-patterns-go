package concrete

import (
	"errors"
	"strings"

	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/handler"
	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/server"
)

// AdminOnly guards the paths under /admin and lets every other request through. It runs after
// Auth, because a role only matters once we know who is calling.
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
