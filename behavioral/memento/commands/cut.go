package commands

import (
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/command"
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/originator"
)

const CutName command.Name = "CUT_SELECTION"

type Cut struct {
	editor originator.Originator
}

func NewCut(editor originator.Originator) *Cut {
	return &Cut{editor: editor}
}

func (c *Cut) Name() command.Name {
	return CutName
}

func (c *Cut) Undoable() bool {
	return true
}

func (c *Cut) Execute() {
	state := c.editor.State()

	if state.SelectionStart == state.SelectionEnd {
		return
	}

	cursor := state.SelectionStart

	state.Clipboard = state.Content[state.SelectionStart:state.SelectionEnd]
	state.Content = state.Content[:state.SelectionStart] + state.Content[state.SelectionEnd:]
	state.Cursor = cursor
	state.SelectionStart = cursor
	state.SelectionEnd = cursor

	c.editor.SetState(state)
}
