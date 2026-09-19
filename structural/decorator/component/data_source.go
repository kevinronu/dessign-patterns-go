package component

type DataSource interface {
	Write(data []byte) error
	Read() ([]byte, error)
}
