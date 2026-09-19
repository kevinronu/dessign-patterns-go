package decorator

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/structural/decorator/component"
)

type Encryption struct {
	wrappee component.DataSource
	aead    cipher.AEAD
}

// NewEncryption returns a decorator that stores authenticated ciphertext. key must be a valid AES key.
func NewEncryption(wrappee component.DataSource, key []byte) (Encryption, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return Encryption{}, fmt.Errorf("new cipher: %w", err)
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return Encryption{}, fmt.Errorf("new gcm: %w", err)
	}

	return Encryption{wrappee: wrappee, aead: aead}, nil
}

func (e Encryption) Write(data []byte) error {
	nonce := make([]byte, e.aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return fmt.Errorf("read nonce: %w", err)
	}

	return e.wrappee.Write(e.aead.Seal(nonce, nonce, data, nil))
}

func (e Encryption) Read() ([]byte, error) {
	data, err := e.wrappee.Read()
	if err != nil {
		return nil, err
	}

	if len(data) < e.aead.NonceSize() {
		return nil, fmt.Errorf("stored blob is %d bytes, shorter than the nonce", len(data))
	}

	nonce, ciphertext := data[:e.aead.NonceSize()], data[e.aead.NonceSize():]

	plain, err := e.aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("open ciphertext: %w", err)
	}

	return plain, nil
}
