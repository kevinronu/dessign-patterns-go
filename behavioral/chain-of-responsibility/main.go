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

	// The head goes on its own line: SetNext returns the next step, not the head of the chain.
	middleware := concrete.NewRateLimit(limit, window)

	// The order is decided here. Over the limit, nothing behind the rate limit runs.
	middleware.SetNext(concrete.NewAuth(srv)).SetNext(concrete.NewAdminOnly(srv))

	srv.SetMiddleware(middleware)

	// One request per reason a step can refuse, plus two that reach the server.
	requests := []handler.Request{
		{Email: "admin@example.com", Password: "admin_pass", Path: "/admin/settings"},
		{Email: "user@example.com", Password: "user_pass", Path: "/reports"},
		{Email: "user@example.com", Password: "user_pass", Path: "/admin/settings"},
		{Email: "admin@example.com", Password: "wrong_pass", Path: "/reports"},
		{Email: "ghost@example.com", Password: "secret", Path: "/reports"},
		{Email: "admin@example.com", Password: "admin_pass", Path: "/reports"},
	}

	fmt.Printf("chain: rate limit (%d per %v) -> auth -> admin only\n\n", limit, window)

	// The client never touches the chain. It learns what came back, and which step refused.
	for _, req := range requests {
		page, err := srv.Serve(req)
		if err != nil {
			fmt.Printf("  %-18s %-11s %-16s %v\n", req.Email, req.Password, req.Path, err)

			continue
		}

		fmt.Printf("  %-18s %-11s %-16s %s\n", req.Email, req.Password, req.Path, page)
	}
}
