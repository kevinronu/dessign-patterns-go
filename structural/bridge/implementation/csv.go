package implementation

type CSV struct{}

func (CSV) Heading(text string) string {
	return "# " + text + "\n"
}

func (CSV) Field(label, value string) string {
	return label + "," + value + "\n"
}
