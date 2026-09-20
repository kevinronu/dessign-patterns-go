package main

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/commands"
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/originator"
)

func main() {
	editor := originator.NewEditor("Hello brave world")
	printState("Initial", editor)

	editor.Execute(commands.NewSelectRange(editor, 0, 5))
	printState("After SELECT_RANGE", editor)

	editor.Execute(commands.NewToggleBold(editor))
	printState("After TOGGLE_BOLD", editor)

	editor.Execute(commands.NewCopy(editor))
	printState("After COPY", editor)

	editor.Execute(commands.NewCut(editor))
	printState("After CUT", editor)

	editor.Execute(commands.NewMoveCaret(editor, len(editor.State().Content)))
	printState("After MOVE_CARET", editor)

	editor.Execute(commands.NewPaste(editor))
	printState("After PASTE", editor)

	undoAndPrint(editor)
	undoAndPrint(editor)

	redoAndPrint(editor)
	redoAndPrint(editor)
	redoAndPrint(editor)
}

func printState(step string, editor *originator.Editor) {
	fmt.Printf("\n%s\n%s\n", step, editor.Debug())
}

func undoAndPrint(editor *originator.Editor) {
	if !editor.Undo() {
		fmt.Println("\nNothing to undo")

		return
	}

	printState("After UNDO", editor)
}

func redoAndPrint(editor *originator.Editor) {
	if !editor.Redo() {
		fmt.Println("\nNothing to redo")

		return
	}

	printState("After REDO", editor)
}
