// Package handler defines the contract of the chain and the link between its steps.
package handler

// Request is what travels down the chain. It carries everything any step might need to decide.
type Request struct {
	Email    string
	Password string
	Path     string
}

// Handler is one step of the chain. Handle returns nil to let the request go on and an error to
// stop it, so the reason travels back to the caller.
//
// SetNext is part of the contract, so a chain can still grow after it starts running.
type Handler interface {
	SetNext(next Handler) Handler
	Handle(req Request) error
}

// Successor is the link to the next step. Concrete handlers embed it, so none of them repeats the
// field, the setter, or the check for the end of the chain.
type Successor struct {
	next Handler
}

// SetNext returns the handler it was given, so the rest of a chain fits in one line. What comes
// back is the next step and not the head, so the caller has to keep the head itself.
func (s *Successor) SetNext(next Handler) Handler {
	s.next = next

	return next
}

// Handle passes the request on, and lets it through when this is the last step. A concrete handler
// hides this method with its own, and calls this one as x.Successor.Handle after its check passed.
func (s *Successor) Handle(req Request) error {
	if s.next == nil {
		return nil
	}

	return s.next.Handle(req)
}
