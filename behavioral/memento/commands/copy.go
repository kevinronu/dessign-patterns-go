package commands

import (
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/command"
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/originator"
)

const CopyName command.Name = "COPY_SELECTION"

type Copy struct {
	editor originator.Originator
}

func NewCopy(editor originator.Originator) *Copy {
	return &Copy{editor: editor}
}

func (c *Copy) Name() command.Name {
	return CopyName
}

func (c *Copy) Undoable() bool {
	return false
}

func (c *Copy) Execute() {
	state := c.editor.State()

	if state.SelectionStart == state.SelectionEnd {
		return
	}

	state.Clipboard = state.Content[state.SelectionStart:state.SelectionEnd]

	c.editor.SetState(state)
}
