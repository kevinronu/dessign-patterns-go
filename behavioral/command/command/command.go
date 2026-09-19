package command

type Command interface {
	Execute() bool
	Undo()
}
