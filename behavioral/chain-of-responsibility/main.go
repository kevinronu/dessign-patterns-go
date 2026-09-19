package main

import (
	"fmt"
	"time"

	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/concrete"
	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/handler"
	"github.com/kevinronu/dessign-patterns-go/behavioral/chain-of-responsibility/server"
)

func main() {
	const (
		limit  = 5
		window = 300 * time.Millisecond
	)

	srv := server.New()
	srv.RegisterAdmin("admin@example.com", "admin_pass")
	srv.Register("user@example.com", "user_pass")

	middleware := concrete.NewRateLimit(limit, window)

	middleware.SetNext(concrete.NewAuth(srv)).SetNext(concrete.NewAdminOnly(srv))

	srv.SetMiddleware(middleware)

	requests := []handler.Request{
		{Email: "admin@example.com", Password: "admin_pass", Path: "/admin/settings"},
		{Email: "user@example.com", Password: "user_pass", Path: "/reports"},
		{Email: "user@example.com", Password: "user_pass", Path: "/admin/settings"},
		{Email: "admin@example.com", Password: "wrong_pass", Path: "/reports"},
		{Email: "ghost@example.com", Password: "secret", Path: "/reports"},
		{Email: "admin@example.com", Password: "admin_pass", Path: "/reports"},
	}

	fmt.Printf("chain: rate limit (%d per %v) -> auth -> admin only\n\n", limit, window)

	for _, req := range requests {
		page, err := srv.Serve(req)
		if err != nil {
			fmt.Printf("  %-18s %-11s %-16s %v\n", req.Email, req.Password, req.Path, err)

			continue
		}

		fmt.Printf("  %-18s %-11s %-16s %s\n", req.Email, req.Password, req.Path, page)
	}
}
