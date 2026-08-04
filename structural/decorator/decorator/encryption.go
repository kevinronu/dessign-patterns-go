package decorator

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"github.com/kevinronu/dessign-patterns-go/structural/decorator/component"
)

// Encryption stores a cipher.AEAD instead of the key, so the key is checked once and then
// dropped.
type Encryption struct {
	wrappee component.DataSource
	aead    cipher.AEAD
}

// NewEncryption returns an error because the key must be 16, 24 or 32 bytes.
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

	// Seal adds the result to its first argument, so passing nonce puts the nonce in front. Read
	// needs that same nonce, and the store keeps only one value.
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
