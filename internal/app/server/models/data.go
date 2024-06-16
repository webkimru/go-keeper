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
	return k.UserID == (ctx.Value("userID")).(int64)
}
