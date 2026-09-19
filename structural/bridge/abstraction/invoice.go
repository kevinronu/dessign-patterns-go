package abstraction

type Invoice struct {
	Document
	Number   string
	Customer string
	Date     string
	Total    string
}

func (i Invoice) Export() string {
	return i.Exporter.Heading("Invoice "+i.Number) +
		i.Exporter.Field("Customer", i.Customer) +
		i.Exporter.Field("Date", i.Date) +
		i.Exporter.Field("Total", i.Total)
}
