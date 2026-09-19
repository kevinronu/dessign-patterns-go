// Package commands provides editor operations as values that can be passed, queued, and invoked.
package commands

import (
	"github.com/kevinronu/dessign-patterns-go/behavioral/command/editor"
)

// Copy copies the editor selection to its clipboard.
type Copy struct {
	commandState
}

func NewCopy(editor *editor.Editor) *Copy {
	return &Copy{commandState: commandState{editor: editor}}
}

func (c *Copy) Execute() bool {
	if c.editor.SelectionStart == c.editor.SelectionEnd {
		return false
	}

	c.editor.Clipboard = c.editor.Text[c.editor.SelectionStart:c.editor.SelectionEnd]

	return false
}
