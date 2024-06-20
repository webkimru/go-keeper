package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewZap(t *testing.T) {
	tests := []struct {
		name    string
		level   string
		wantErr bool
	}{
		{"positive", "info", false},
		{"negative", "none", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewZap(tt.level)
			if (err != nil) != tt.wantErr {
				assert.NoError(t, err)
			}
		})
	}
}
