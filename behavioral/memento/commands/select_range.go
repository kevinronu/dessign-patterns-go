package commands

import (
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/command"
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/originator"
)

const SelectRangeName command.Name = "SELECT_RANGE"

type SelectRange struct {
	editor originator.Originator
	start  int
	end    int
}

func NewSelectRange(editor originator.Originator, start, end int) *SelectRange {
	return &SelectRange{editor: editor, start: start, end: end}
}

func (s *SelectRange) Name() command.Name {
	return SelectRangeName
}

func (s *SelectRange) Undoable() bool {
	return false
}

func (s *SelectRange) Execute() {
	state := s.editor.State()

	start := min(max(0, s.start), len(state.Content))
	end := min(max(start, s.end), len(state.Content))

	state.SelectionStart = start
	state.SelectionEnd = end

	s.editor.SetState(state)
}
