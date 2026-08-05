// Package server holds the service the chain protects. It does its own work, and it takes one
// middleware to run before that work starts.
package server

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/handler"
)

type account struct {
	password string
	admin    bool
}

// pages is the work the server does once the middleware lets a request through.
var pages = map[string]string{
	"/reports":        "quarterly report",
	"/admin/settings": "server settings",
}

// Server keeps the accounts and one middleware. The middleware is a handler.Handler, so a whole
// chain fits in that single field and the server never learns how long it is.
type Server struct {
	accounts   map[string]account
	middleware handler.Handler
}

func New() *Server {
	return &Server{accounts: make(map[string]account)}
}

// SetMiddleware puts a chain in front of every request. Without it the server answers on its own.
func (s *Server) SetMiddleware(middleware handler.Handler) {
	s.middleware = middleware
}

func (s *Server) Register(email, password string) {
	s.accounts[email] = account{password: password}
}

func (s *Server) RegisterAdmin(email, password string) {
	s.accounts[email] = account{password: password, admin: true}
}

// Serve runs the middleware first and only then does the work, so the error of the step that
// refused comes back unchanged.
func (s *Server) Serve(req handler.Request) (string, error) {
	if s.middleware != nil {
		if err := s.middleware.Handle(req); err != nil {
			return "", err
		}
	}

	page, ok := pages[req.Path]
	if !ok {
		return "", fmt.Errorf("no page at %q", req.Path)
	}

	return page, nil
}

func (s *Server) HasEmail(email string) bool {
	_, ok := s.accounts[email]

	return ok
}

func (s *Server) IsValidPassword(email, password string) bool {
	return s.accounts[email].password == password
}

func (s *Server) IsAdmin(email string) bool {
	return s.accounts[email].admin
}
