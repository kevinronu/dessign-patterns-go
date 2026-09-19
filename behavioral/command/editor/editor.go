package editor

import "github.com/kevinronu/dessign-patterns-go/behavioral/command/command"

type Editor struct {
	Text           string
	Clipboard      string
	SelectionStart int
	SelectionEnd   int
	CaretPosition  int
	history        command.History
}

func New(text string) *Editor {
	return &Editor{
		Text:           text,
		SelectionStart: len(text),
		SelectionEnd:   len(text),
		CaretPosition:  len(text),
	}
}

func (e *Editor) Execute(cmd command.Command) {
	if cmd.Execute() {
		e.history.Push(cmd)
	}
}

func (e *Editor) Undo() {
	if e.history.Empty() {
		return
	}

	e.history.Pop().Undo()
}
