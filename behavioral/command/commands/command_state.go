package commands

import "github.com/kevinronu/dessign-patterns-go/behavioral/command/editor"

type commandState struct {
	editor     *editor.Editor
	backupText string
}

func (c *commandState) backup() {
	c.backupText = c.editor.Text
}

func (c *commandState) Undo() {
	c.editor.Text = c.backupText
}
