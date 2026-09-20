package originator

import (
	"fmt"
	"strings"

	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/caretaker"
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/command"
)

type Editor struct {
	state   State
	history caretaker.History
}

func NewEditor(initialText string) *Editor {
	return &Editor{state: State{
		Content:        initialText,
		Cursor:         len(initialText),
		SelectionStart: len(initialText),
		SelectionEnd:   len(initialText),
	}}
}

// The checkpoint is pushed before the command runs, because the memento has to
// capture the state the command is about to replace.
func (e *Editor) Execute(command command.Command) {
	if command.Undoable() {
		e.history.Push(caretaker.Checkpoint{
			Command: command,
			Memento: NewEditorMemento(e),
		})
	}

	command.Execute()
}

// Undo reports false when the history is exhausted, so a caller can repaint its
// view only when something changed.
func (e *Editor) Undo() bool {
	return e.history.Undo()
}

// Redo reports false when there is nothing left to reapply.
func (e *Editor) Redo() bool {
	return e.history.Redo()
}

func (e *Editor) State() State {
	return e.state.Clone()
}

func (e *Editor) SetState(state State) {
	e.state = state.Clone()
}

func (e *Editor) Debug() string {
	return strings.Join([]string{
		"— Editor State —",
		fmt.Sprintf("content    : %q", e.state.Content),
		fmt.Sprintf("cursor     : %d", e.state.Cursor),
		fmt.Sprintf("selection  : [%d, %d]", e.state.SelectionStart, e.state.SelectionEnd),
		fmt.Sprintf("clipboard  : %q", e.state.Clipboard),
		fmt.Sprintf("bold       : %v", e.state.Bold),
	}, "\n")
}
