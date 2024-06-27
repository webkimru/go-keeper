package models

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyValue_Validate(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		required []string
		want     string
		wantErr  bool
	}{
		{"positive: title", "title", []string{"title"}, "title", false},
		{"negative: title", "", []string{"title"}, "title", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KeyValue{Title: tt.field}
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

func TestKeyValue_valid(t *testing.T) {
	tests := []struct {
		name  string
		field string
		arg   string
		want  bool
	}{
		{"positive", "title", "title", true},
		{"negative: id", "", "id", false},
		{"negative: title", "", "title", false},
		{"negative: key", "", "key", false},
		{"negative: value", "", "value", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KeyValue{Title: tt.field}
			assert.Equalf(t, tt.want, k.valid(tt.arg), "valid(%v)", tt.arg)
		})
	}
}

func TestKeyValue_getContextUserID(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		ctx  context.Context
		want int64
	}{
		{"positive", context.WithValue(ctx, ContextKey("userID"), int64(1)), 1},
		{"negative", context.WithValue(ctx, ContextKey("userID"), 1), -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KeyValue{}
			assert.Equalf(t, tt.want, k.getContextUserID(tt.ctx), "getContextUserID(%v)", tt.ctx)
		})
	}
}
