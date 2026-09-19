package commands

import "github.com/kevinronu/dessign-patterns-go/behavioral/command/editor"

type Paste struct {
	commandState
}

func NewPaste(editor *editor.Editor) *Paste {
	return &Paste{commandState: commandState{editor: editor}}
}

func (p *Paste) Execute() bool {
	if p.editor.Clipboard == "" {
		return false
	}

	p.backup()
	p.editor.Text = p.editor.Text[:p.editor.CaretPosition] + p.editor.Clipboard + p.editor.Text[p.editor.CaretPosition:]

	newCaretPosition := p.editor.CaretPosition + len(p.editor.Clipboard)
	p.editor.CaretPosition = newCaretPosition
	p.editor.SelectionStart = newCaretPosition
	p.editor.SelectionEnd = newCaretPosition

	return true
}
