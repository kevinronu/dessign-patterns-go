// Package service defines the contract the client depends on. The real service and the proxy both
// implement it, so the client cannot tell which one it holds.
package service

type VideoLibrary interface {
	Download(id string) ([]byte, error)
}
