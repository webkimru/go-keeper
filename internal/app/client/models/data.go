package models

import (
	"context"
	"fmt"
)

const (
	KeyValueStateNew       = "NEW"
	KeyValueStateProcessed = "PROCESSED"
)

// KeyValue contains key-value entity information.
type KeyValue struct {
	ID        int64
	UserID    int64
	Title     string
	Key       string
	Value     string
	Status    string
	CreatedAt string
	UpdatedAt string
}

// Validate is a wrapper over valid() to decorate errors.
func (k *KeyValue) Validate(required ...string) (string, error) {
	for _, field := range required {
		if !k.valid(field) {
			return field, fmt.Errorf("field %s is required", field)
		}
	}

	return "", nil
}

// valid checks entering data fields.
func (k *KeyValue) valid(field string) bool {
	switch field {
	case "id":
		if k.ID == 0 {
			return false
		}
	case "user_id":
		if k.UserID == 0 {
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

// CanAccess compares a user ID in an entering model with context user ID.
func (k *KeyValue) CanAccess(ctx context.Context) bool {
	return k.UserID == k.getContextUserID(ctx)
}

// ContextKey is a custom type for context usage.
type ContextKey string

// getContextUserID is a helper method to get user ID out of ctx.Value as int64.
func (k *KeyValue) getContextUserID(ctx context.Context) int64 {
	id := ctx.Value(ContextKey("userID"))

	switch id := id.(type) {
	case int64:
		return id

	default:
		return -1
	}
}

// GetContextUserID public method wrapped getContextUserID.
func (k *KeyValue) GetContextUserID(ctx context.Context) int64 {
	return k.getContextUserID(ctx)
}
