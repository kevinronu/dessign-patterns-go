package abstraction

// Report is another refined abstraction: different content, composed with the same primitives, in any format.
type Report struct {
	Document
	Title   string
	Author  string
	Date    string
	Summary string
}

func (r Report) Export() string {
	return r.Exporter.Heading(r.Title) +
		r.Exporter.Field("Author", r.Author) +
		r.Exporter.Field("Date", r.Date) +
		r.Exporter.Field("Summary", r.Summary)
}
