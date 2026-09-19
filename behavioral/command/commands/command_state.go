package commands

import "github.com/kevinronu/dessign-patterns-go/behavioral/command/editor"

// commandState holds the receiver and backup shared by reversible commands.
type commandState struct {
	editor     *editor.Editor
	backupText string
}

func (c *commandState) backup() {
	c.backupText = c.editor.Text
}

// Undo restores the text that backup saved before a command ran.
func (c *commandState) Undo() {
	c.editor.Text = c.backupText
}
