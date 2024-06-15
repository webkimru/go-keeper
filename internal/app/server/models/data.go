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

func (k *KeyValue) Validate() (string, error) {
	if err := k.required(k.Key); err != nil {
		return "key", err
	}
	if err := k.required(k.Value); err != nil {
		return "value", err
	}

	return "", nil
}

func (k *KeyValue) required(field string) error {
	if field == "" {
		return fmt.Errorf("field %s is required", field)
	}

	return nil
}

func (k *KeyValue) CanAccess(ctx context.Context) bool {
	return k.UserID == (ctx.Value("userID")).(int64)
}
