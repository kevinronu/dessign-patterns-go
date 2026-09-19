package commands

import (
	"github.com/kevinronu/dessign-patterns-go/behavioral/command/editor"
)

// Cut removes the editor selection and keeps its prior text for Undo.
type Cut struct {
	editor *editor.Editor
	backup string
}

func NewCut(editor *editor.Editor) *Cut {
	return &Cut{editor: editor}
}

func (c *Cut) Execute() bool {
	if c.editor.SelectionStart == c.editor.SelectionEnd {
		return false
	}

	c.backup = c.editor.Text
	c.editor.Clipboard = c.editor.Text[c.editor.SelectionStart:c.editor.SelectionEnd]
	c.editor.Text = c.editor.Text[:c.editor.SelectionStart] + c.editor.Text[c.editor.SelectionEnd:]
	c.editor.SelectionEnd = c.editor.SelectionStart
	c.editor.CaretPosition = c.editor.SelectionStart

	return true
}

func (c *Cut) Undo() {
	c.editor.Text = c.backup
}
