package command

// History keeps executed commands in last-in, first-out order.
type History struct {
	commands []Command
}

func (h *History) Push(command Command) {
	h.commands = append(h.commands, command)
}

// Pop returns the most recent command, or nil when the history is empty.
func (h *History) Pop() Command {
	if h.Empty() {
		return nil
	}

	last := len(h.commands) - 1
	command := h.commands[last]
	h.commands = h.commands[:last]

	return command
}

func (h History) Empty() bool {
	return len(h.commands) == 0
}
