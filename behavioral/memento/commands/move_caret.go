package commands

import (
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/command"
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/originator"
)

const MoveCaretName command.Name = "MOVE_CARET"

type MoveCaret struct {
	editor originator.Originator
	target int
}

func NewMoveCaret(editor originator.Originator, target int) *MoveCaret {
	return &MoveCaret{editor: editor, target: target}
}

func (m *MoveCaret) Name() command.Name {
	return MoveCaretName
}

func (m *MoveCaret) Undoable() bool {
	return false
}

func (m *MoveCaret) Execute() {
	state := m.editor.State()

	position := min(max(0, m.target), len(state.Content))

	state.Cursor = position
	state.SelectionStart = position
	state.SelectionEnd = position

	m.editor.SetState(state)
}
