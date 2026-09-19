package commands

import "github.com/kevinronu/dessign-patterns-go/behavioral/command/editor"

type Cut struct {
	commandState
}

func NewCut(editor *editor.Editor) *Cut {
	return &Cut{commandState: commandState{editor: editor}}
}

func (c *Cut) Execute() bool {
	if c.editor.SelectionStart == c.editor.SelectionEnd {
		return false
	}

	c.backup()
	c.editor.Clipboard = c.editor.Text[c.editor.SelectionStart:c.editor.SelectionEnd]
	c.editor.Text = c.editor.Text[:c.editor.SelectionStart] + c.editor.Text[c.editor.SelectionEnd:]
	c.editor.SelectionEnd = c.editor.SelectionStart
	c.editor.CaretPosition = c.editor.SelectionStart

	return true
}
