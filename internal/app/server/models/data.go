package models

import "fmt"

type KeyValue struct {
	ID        int64
	UserID    int64
	Title     string
	Key       string
	Value     string
	CreatedAt string
}

func (k *KeyValue) Validate() error {
	if err := k.required(k.Key); err != nil {
		return err
	}
	if err := k.required(k.Value); err != nil {
		return err
	}

	return nil
}

func (k *KeyValue) required(field string) error {
	if field == "" {
		return fmt.Errorf("field %s is required", field)
	}

	return nil
}
