// Package component defines the contract shared by the store and its decorators.
package component

// DataSource is all the client sees, so a plain store and a full stack can replace each other.
type DataSource interface {
	Write(data []byte) error
	Read() ([]byte, error)
}
