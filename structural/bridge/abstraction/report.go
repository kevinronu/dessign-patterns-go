package abstraction

// Report shows what the split buys: a second document type needs no change to any Exporter.
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
