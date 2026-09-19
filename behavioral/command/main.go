package main

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/behavioral/command/commands"
	"github.com/kevinronu/dessign-patterns-go/behavioral/command/editor"
)

func main() {
	document := editor.New("Hello, Command!")
	document.SelectionStart = 7
	document.SelectionEnd = 14

	document.Execute(commands.NewCut(document))
	fmt.Printf("cut:   text=%q clipboard=%q\n", document.Text, document.Clipboard)

	document.Execute(commands.NewPaste(document))
	fmt.Printf("paste: text=%q\n", document.Text)

	document.Undo()
	fmt.Printf("undo:  text=%q\n", document.Text)

	document.Undo()
	fmt.Printf("undo:  text=%q clipboard=%q\n", document.Text, document.Clipboard)
}
