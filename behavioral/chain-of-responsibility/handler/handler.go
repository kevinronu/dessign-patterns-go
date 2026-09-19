package handler

type Request struct {
	Email    string
	Password string
	Path     string
}

type Handler interface {
	SetNext(next Handler) Handler
	Handle(req Request) error
}

type Successor struct {
	next Handler
}

func (s *Successor) SetNext(next Handler) Handler {
	s.next = next

	return next
}

func (s *Successor) Handle(req Request) error {
	if s.next == nil {
		return nil
	}

	return s.next.Handle(req)
}
