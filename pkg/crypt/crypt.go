// Package crypt implements encryption and decryption
package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
)

// Crypt is a structure fot encryption and decryption
type Crypt struct {
	Aesgcm cipher.AEAD
	Nonce  []byte
}

// New returns a new encryption.
func New(secret string) (*Crypt, error) {
	aesgcm, nonce, err := Initialize(secret)
	if err != nil {
		return nil, err
	}

	return &Crypt{Aesgcm: aesgcm, Nonce: nonce}, nil
}

// Encrypt encrypt data.
func (c *Crypt) Encrypt(str string) string {
	encrypted := c.Aesgcm.Seal(nil, c.Nonce, []byte(str), nil)

	return hex.EncodeToString(encrypted)
}

// Decrypt decrypt data.
func (c *Crypt) Decrypt(str string) (string, error) {

	decoded, err := hex.DecodeString(str)
	if err != nil {
		return "", err
	}

	decrypted, err := c.Aesgcm.Open(nil, c.Nonce, decoded, nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

// Initialize prepares GCM and initialization vector
func Initialize(secret string) (cipher.AEAD, []byte, error) {
	key := sha256.Sum256([]byte(secret))

	// NewCipher creates and returns a new cipher.Block.
	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, nil, err
	}

	// NewGCM returns the given 128-bit, block cipher wrapped in Galois Counter Mode (GCM) with the standard nonce length.
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, nil, err
	}

	// Create an initialization vector.
	nonce := key[len(key)-aesgcm.NonceSize():]

	return aesgcm, nonce, nil
}
