package implementation

type HTML struct{}

func (HTML) Heading(text string) string {
	return "<h1>" + text + "</h1>\n"
}

func (HTML) Field(label, value string) string {
	return "<p><b>" + label + ":</b> " + value + "</p>\n"
}
