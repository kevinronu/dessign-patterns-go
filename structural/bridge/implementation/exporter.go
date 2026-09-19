package implementation

type Exporter interface {
	Heading(text string) string
	Field(label, value string) string
}
