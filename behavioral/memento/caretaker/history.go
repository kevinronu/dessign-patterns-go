package caretaker

import (
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/command"
	"github.com/kevinronu/dessign-patterns-go/behavioral/memento/memento"
)

type Checkpoint struct {
	Command command.Command
	Memento memento.Memento
}

// History keeps undone checkpoints instead of discarding them, so they stay
// redoable.
type History struct {
	checkpoints []Checkpoint
	applied     int
}

func (h *History) Push(checkpoint Checkpoint) {
	if h.applied != len(h.checkpoints) && h.applied > 0 {
		h.checkpoints = h.checkpoints[:h.applied]
	}

	h.checkpoints = append(h.checkpoints, checkpoint)
	h.applied = len(h.checkpoints)
}

func (h *History) Undo() bool {
	undone, ok := h.undoCheckpoint()
	if !ok {
		return false
	}

	fmt.Println("Undoing:", undone.Command.Name())

	undone.Memento.Restore()

	return true
}

func (h *History) undoCheckpoint() (Checkpoint, bool) {
	if h.applied == 0 {
		return Checkpoint{}, false
	}

	h.applied--

	return h.checkpoints[h.applied], true
}

func (h *History) Redo() bool {
	redone, ok := h.redoCheckpoint()
	if !ok {
		return false
	}

	fmt.Println("Redoing:", redone.Command.Name())

	redone.Memento.Restore()
	redone.Command.Execute()

	return true
}

func (h *History) redoCheckpoint() (Checkpoint, bool) {
	if h.applied == len(h.checkpoints) {
		return Checkpoint{}, false
	}

	h.applied++

	return h.checkpoints[h.applied-1], true
}
