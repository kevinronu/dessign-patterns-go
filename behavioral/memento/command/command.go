package command

type Command interface {
	Name() Name
	Execute()
	// Undoable reports whether the command mutates the document, which is what
	// an editor history records; caret and selection moves are not undone.
	Undoable() bool
}

type Name string
