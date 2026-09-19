package service

type VideoLibrary interface {
	Download(id string) ([]byte, error)
}
