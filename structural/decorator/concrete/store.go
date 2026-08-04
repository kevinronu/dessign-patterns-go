// Package concrete holds the concrete component: the DataSource that decorators wrap, and the
// only one that adds no behavior of its own.
package concrete

// Store keeps the bytes in an exported field so the demo can see what really reached it.
type Store struct {
	Data []byte
}

// Write cannot fail here. The contract keeps the error because a real store can fail.
func (s *Store) Write(data []byte) error {
	s.Data = data

	return nil
}

func (s *Store) Read() ([]byte, error) {
	return s.Data, nil
}
