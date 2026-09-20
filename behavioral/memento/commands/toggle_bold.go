package commands

import (
	"slices"

	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/command"
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/originator"
)

const ToggleBoldName command.Name = "TOGGLE_BOLD"

type ToggleBold struct {
	editor originator.Originator
}

func NewToggleBold(editor originator.Originator) *ToggleBold {
	return &ToggleBold{editor: editor}
}

func (t *ToggleBold) Name() command.Name {
	return ToggleBoldName
}

func (t *ToggleBold) Undoable() bool {
	return true
}

func (t *ToggleBold) Execute() {
	state := t.editor.State()

	if state.SelectionStart == state.SelectionEnd {
		return
	}

	selection := originator.Range{Start: state.SelectionStart, End: state.SelectionEnd}

	if bolded := slices.Index(state.Bold, selection); bolded >= 0 {
		state.Bold = slices.Delete(state.Bold, bolded, bolded+1)
	} else {
		state.Bold = append(state.Bold, selection)
	}

	t.editor.SetState(state)
}
