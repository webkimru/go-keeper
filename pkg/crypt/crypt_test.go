package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrypt_Decrypt(t *testing.T) {

	key := sha256.Sum256([]byte("secret"))
	aesblock, _ := aes.NewCipher(key[:])
	aesgcm, _ := cipher.NewGCM(aesblock)

	nonce := key[len(key)-aesgcm.NonceSize():]
	nonceWrong, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda")

	type fields struct {
		Aesgcm cipher.AEAD
		Nonce  []byte
	}
	tests := []struct {
		name    string
		fields  fields
		arg     string
		want    string
		wantErr bool
	}{
		{
			"positive: correct data",
			fields{aesgcm, nonce},
			"df0c76be5b07aee90dd132c0103722ebb99c60c6a9",
			"admin",
			false,
		},
		{
			"negative: wrong nonce",
			fields{aesgcm, nonceWrong},
			"df0c76be5b07aee90dd132c0103722ebb99c60c6a9",
			"",
			true,
		},
		{
			"negative: empty payload",
			fields{aesgcm, nonce},
			"wrong hex",
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Crypt{
				Aesgcm: tt.fields.Aesgcm,
				Nonce:  tt.fields.Nonce,
			}
			got, err := c.Decrypt(tt.arg)
			if (err != nil) != tt.wantErr {
				assert.NoError(t, err)
				return
			}
			if got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestInitialize(t *testing.T) {
	tests := []struct {
		name    string
		secret  string
		wantErr bool
	}{
		{"positive", "secret", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := Initialize(tt.secret)
			if (err != nil) != tt.wantErr {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		secret  string
		wantErr bool
	}{
		{"positive", "secret", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.secret)
			if (err != nil) != tt.wantErr {
				assert.NoError(t, err)
			}
		})
	}
}
