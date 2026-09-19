package concrete

type Store struct {
	Data []byte
}

func (s *Store) Write(data []byte) error {
	s.Data = data

	return nil
}

// Read returns the store's backing slice. Callers must not modify it.
func (s *Store) Read() ([]byte, error) {
	return s.Data, nil
}
