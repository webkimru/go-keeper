package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		required []string
		want     string
		wantErr  bool
	}{
		{"positive: login", "login", []string{"login"}, "login", false},
		{"negative: login", "", []string{"login"}, "login", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &User{Login: tt.field}
			got, err := k.Validate(tt.required...)
			if (err != nil) != tt.wantErr {
				assert.NoError(t, err)
				return
			}
			if got != tt.want {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUser_valid(t *testing.T) {
	tests := []struct {
		name  string
		field string
		arg   string
		want  bool
	}{
		{"positive", "login", "login", true},
		{"negative: login", "", "login", false},
		{"negative: password", "", "password", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &User{Login: tt.field}
			assert.Equalf(t, tt.want, k.valid(tt.arg), "valid(%v)", tt.arg)
		})
	}
}
