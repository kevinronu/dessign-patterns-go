// Package command defines the command contract and its undo history.
package command

// Command represents one operation that may be undone.
type Command interface {
	Execute() bool
	Undo()
}
