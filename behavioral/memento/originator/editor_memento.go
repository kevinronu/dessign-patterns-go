package originator

type EditorMemento struct {
	editor Originator
	state  State
}

func NewEditorMemento(editor Originator) *EditorMemento {
	return &EditorMemento{editor: editor, state: editor.State()}
}

func (m *EditorMemento) Restore() {
	m.editor.SetState(m.state)
}
