package commands

import (
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/command"
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/originator"
)

const PasteName command.Name = "PASTE_FROM_CLIPBOARD"

type Paste struct {
	editor originator.Originator
}

func NewPaste(editor originator.Originator) *Paste {
	return &Paste{editor: editor}
}

func (p *Paste) Name() command.Name {
	return PasteName
}

func (p *Paste) Undoable() bool {
	return true
}

func (p *Paste) Execute() {
	state := p.editor.State()

	if state.Clipboard == "" {
		return
	}

	cursor := state.Cursor + len(state.Clipboard)

	state.Content = state.Content[:state.Cursor] + state.Clipboard + state.Content[state.Cursor:]
	state.Cursor = cursor
	state.SelectionStart = cursor
	state.SelectionEnd = cursor

	p.editor.SetState(state)
}
