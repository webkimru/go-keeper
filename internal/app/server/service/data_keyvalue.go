package service

import (
	"context"
	"fmt"
	"github.com/webkimru/go-keeper/pkg/errs"

	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/crypt"
)

// KeyValueStore is an interface to store data.
type KeyValueStore interface {
	Add(ctx context.Context, model models.KeyValue) error
	Get(ctx context.Context, id int64) (*models.KeyValue, error)
	List(ctx context.Context) ([]models.KeyValue, error)
	Update(ctx context.Context, model models.KeyValue) error
	Delete(ctx context.Context, id int64) error
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
func (s *KeyValueService) List(ctx context.Context) ([]models.KeyValue, error) {
	data, err := s.storage.List(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
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
func (s *KeyValueService) Delete(ctx context.Context, id int64) error {
	if err := s.storage.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *KeyValueService) Decrypt(field string) (string, error) {
	decrypted, err := s.cryptManager.Decrypt(field)
	if err != nil {
		return "", fmt.Errorf("KeyValueService - Get - s.cryptManager.Decrypt(%s): %w", field, err)
	}

	return decrypted, nil
}
