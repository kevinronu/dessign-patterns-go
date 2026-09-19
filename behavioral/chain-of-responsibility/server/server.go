package server

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/handler"
)

type account struct {
	password string
	admin    bool
}

var pages = map[string]string{
	"/reports":        "quarterly report",
	"/admin/settings": "server settings",
}

type Server struct {
	accounts   map[string]account
	middleware handler.Handler
}

func New() *Server {
	return &Server{accounts: make(map[string]account)}
}

func (s *Server) SetMiddleware(middleware handler.Handler) {
	s.middleware = middleware
}

func (s *Server) Register(email, password string) {
	s.accounts[email] = account{password: password}
}

func (s *Server) RegisterAdmin(email, password string) {
	s.accounts[email] = account{password: password, admin: true}
}

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
