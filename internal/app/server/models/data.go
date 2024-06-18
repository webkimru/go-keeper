package models

import (
	"context"
	"fmt"
)

type KeyValue struct {
	ID        int64
	UserID    int64
	Title     string
	Key       string
	Value     string
	CreatedAt string
	UpdatedAt string
}

func (k *KeyValue) Validate(required ...string) (string, error) {
	for _, field := range required {
		if !k.valid(field) {
			return field, fmt.Errorf("field %s is required", field)
		}
	}

	return "", nil
}

func (k *KeyValue) valid(field string) bool {
	switch field {
	case "id":
		if k.ID == 0 {
			return false
		}
	case "title":
		if k.Title == "" {
			return false
		}
	case "key":
		if k.Key == "" {
			return false
		}
	case "value":
		if k.Value == "" {
			return false
		}
	}

	return true
}

func (k *KeyValue) CanAccess(ctx context.Context) bool {
	return k.UserID == k.getContextUserID(ctx)
}

type ContextKey string

func (k *KeyValue) getContextUserID(ctx context.Context) int64 {
	id := ctx.Value(ContextKey("userID"))

	switch id := id.(type) {
	case int64:
		return id

	default:
		return -1
	}
}
