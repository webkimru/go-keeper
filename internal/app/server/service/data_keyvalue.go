package service

import (
	"context"
	"fmt"

	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/crypt"
	"github.com/webkimru/go-keeper/pkg/errs"
)

//go:generate mockgen -destination=mocks/mock_keyvalue.go -package=mocks github.com/webkimru/go-keeper/internal/app/server/service KeyValueStore

// KeyValueStore is an interface to store data.
type KeyValueStore interface {
	Add(ctx context.Context, model models.KeyValue) error
	Get(ctx context.Context, id int64) (*models.KeyValue, error)
	List(ctx context.Context, userID, limit, offset int64) ([]models.KeyValue, error)
	Update(ctx context.Context, model models.KeyValue) error
	Delete(ctx context.Context, userID, id int64) error
}

// KeyValueService contains data storage.
type KeyValueService struct {
	storage      KeyValueStore
	cryptManager *crypt.Crypt
}

// NewKeyValueService return a new service with storage.
func NewKeyValueService(storage KeyValueStore, cryptManager *crypt.Crypt) *KeyValueService {
	return &KeyValueService{storage: storage, cryptManager: cryptManager}
}

// Add puts data to the storage.
func (s *KeyValueService) Add(ctx context.Context, model models.KeyValue) error {
	if field, err := model.Validate("title", "key", "value"); err != nil {
		return fmt.Errorf("%s: %w", field, errs.ErrBadRequest)
	}

	// encrypt
	model.Key = s.cryptManager.Encrypt(model.Key)
	model.Value = s.cryptManager.Encrypt(model.Value)

	if err := s.storage.Add(ctx, model); err != nil {
		return err
	}

	return nil
}

// Get returns a row of the data.
func (s *KeyValueService) Get(ctx context.Context, id int64) (*models.KeyValue, error) {
	data, err := s.storage.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// check access
	if !data.CanAccess(ctx) {
		return nil, errs.ErrPermissionDenied
	}

	// decrypt
	if data.Key, err = s.Decrypt(data.Key); err != nil {
		return nil, err
	}
	if data.Value, err = s.Decrypt(data.Value); err != nil {
		return nil, err
	}

	return data, nil
}

// List returns a slice of the data.
func (s *KeyValueService) List(ctx context.Context, userID, limit, offset int64) ([]models.KeyValue, error) {
	if limit == 0 {
		return nil, fmt.Errorf("%s: %w", "limit > 0", errs.ErrBadRequest)
	}

	data, err := s.storage.List(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	// decrypt
	slice := make([]models.KeyValue, len(data))
	for _, item := range data {
		if item.Key, err = s.Decrypt(item.Key); err != nil {
			return nil, err
		}
		if item.Value, err = s.Decrypt(item.Value); err != nil {
			return nil, err
		}
		slice = append(slice, models.KeyValue{
			Key:   item.Key,
			Value: item.Value,
		})
	}

	return slice, nil
}

// Update updates a row of the data.
func (s *KeyValueService) Update(ctx context.Context, model models.KeyValue) error {
	if field, err := model.Validate("id", "title", "key", "value"); err != nil {
		return fmt.Errorf("%s: %w", field, errs.ErrBadRequest)
	}

	// encrypt
	model.Key = s.cryptManager.Encrypt(model.Key)
	model.Value = s.cryptManager.Encrypt(model.Value)

	if err := s.storage.Update(ctx, model); err != nil {
		return err
	}

	return nil
}

// Delete deletes a row of the data.
func (s *KeyValueService) Delete(ctx context.Context, userID, id int64) error {
	if err := s.storage.Delete(ctx, userID, id); err != nil {
		return err
	}

	return nil
}

// Decrypt decrypts fields.
func (s *KeyValueService) Decrypt(field string) (string, error) {
	decrypted, err := s.cryptManager.Decrypt(field)
	if err != nil {
		return "", fmt.Errorf("KeyValueService - Get - s.cryptManager.Decrypt(%s): %w", field, err)
	}

	return decrypted, nil
}
