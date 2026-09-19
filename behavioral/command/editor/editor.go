// Package editor contains the command receiver and its undo history.
package editor

import "github.com/kevinronu/dessign-patterns-go/behavioral/command/command"

// Editor is the receiver that commands change.
type Editor struct {
	Text           string
	Clipboard      string
	SelectionStart int
	SelectionEnd   int
	CaretPosition  int
	history        command.History
}

// New returns an editor with its caret at the end of text.
func New(text string) *Editor {
	return &Editor{
		Text:           text,
		SelectionStart: len(text),
		SelectionEnd:   len(text),
		CaretPosition:  len(text),
	}
}

// Execute records a command only when it changes the editor.
func (e *Editor) Execute(cmd command.Command) {
	if cmd.Execute() {
		e.history.Push(cmd)
	}
}

// Undo reverses the most recently executed command.
func (e *Editor) Undo() {
	if e.history.Empty() {
		return
	}

	e.history.Pop().Undo()
}
