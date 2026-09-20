package originator

import "slices"

type Range struct {
	Start int
	End   int
}

type State struct {
	Content        string
	Cursor         int
	SelectionStart int
	SelectionEnd   int
	Clipboard      string
	Bold           []Range
}

// Clone deep copies the reference fields. Assigning a State only copies the slice
// header, so a snapshot would keep sharing its elements with the live state.
func (s State) Clone() State {
	s.Bold = slices.Clone(s.Bold)

	return s
}
