package service

import (
	"context"
	"fmt"

	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/pkg/crypt"
	"github.com/webkimru/go-keeper/pkg/errs"
	"github.com/webkimru/go-keeper/pkg/logger"
)

const (
	_defaultRecordsLimit  = 100
	_defaultRecordsOffset = 0
)

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
	logger       *logger.Log
	cryptManager *crypt.Crypt
	storage      KeyValueStore
}

// NewKeyValueService return a new service with storage.
func NewKeyValueService(storage KeyValueStore, cryptManager *crypt.Crypt, l *logger.Log) *KeyValueService {
	return &KeyValueService{storage: storage, cryptManager: cryptManager, logger: l}
}

// Add puts data to the storage.
func (s *KeyValueService) Add(ctx context.Context, model models.KeyValue) error {
	userID := model.GetContextUserID(ctx)
	if userID == -1 {
		return fmt.Errorf("KeyValueService - Add - model.GetContextUserID(): %w", errs.ErrPermissionDenied)
	}

	if field, err := model.Validate("title", "key", "value"); err != nil {
		return fmt.Errorf("KeyValueService - Add - model.Validate(): %w: %s is required", errs.ErrBadRequest, field)
	}

	model.UserID = userID
	model.Status = models.KeyValueStateNew
	// encrypt
	model.Key = s.cryptManager.Encrypt(model.Key)
	model.Value = s.cryptManager.Encrypt(model.Value)

	if err := s.storage.Add(ctx, model); err != nil {
		return fmt.Errorf("KeyValueService - Add - s.storage.Add(): %w", err)
	}

	return nil
}

// Get returns a row of the data.
func (s *KeyValueService) Get(ctx context.Context, id int64) (*models.KeyValue, error) {
	// validate
	model := models.KeyValue{ID: id}
	if field, err := model.Validate("id"); err != nil {
		return nil, fmt.Errorf("KeyValueService - Get - model.Validate(): %w: %s is required", errs.ErrBadRequest, field)
	}

	data, err := s.storage.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("KeyValueService - Get - s.storage.Get(): %w", err)
	}

	// check access
	if !data.CanAccess(ctx) {
		return nil, fmt.Errorf("KeyValueService - Get - data.CanAccess(): %w", errs.ErrPermissionDenied)
	}

	// decrypt
	if data.Key, err = s.Decrypt(data.Key); err != nil {
		return nil, fmt.Errorf("KeyValueService - Get - s.Decrypt(data.Key): %w", err)
	}
	if data.Value, err = s.Decrypt(data.Value); err != nil {
		return nil, fmt.Errorf("KeyValueService - Get - s.Decrypt(data.Value): %w", err)
	}

	return data, nil
}

// List returns a slice of the data.
func (s *KeyValueService) List(ctx context.Context) ([]models.KeyValue, error) {
	model := models.KeyValue{}
	userID := model.GetContextUserID(ctx)
	if userID == -1 {
		return nil, fmt.Errorf("KeyValueService - Add - model.GetContextUserID(): %w", errs.ErrPermissionDenied)
	}
	model.UserID = userID
	// validate
	if field, err := model.Validate("user_id"); err != nil {
		return nil, fmt.Errorf("KeyValueService - List - model.Validate(): %w: %s is required", errs.ErrBadRequest, field)
	}

	data, err := s.storage.List(ctx, userID, _defaultRecordsLimit, _defaultRecordsOffset)
	if err != nil {
		return nil, fmt.Errorf("KeyValueService - List - s.storage.List(): %w", err)
	}

	// decrypt
	slice := make([]models.KeyValue, len(data))
	n := 0
	for _, item := range data {
		if item.Key, err = s.Decrypt(item.Key); err != nil {
			return nil, fmt.Errorf("KeyValueService - List - s.Decrypt(item.Key): %w", err)
		}
		if item.Value, err = s.Decrypt(item.Value); err != nil {
			return nil, fmt.Errorf("KeyValueService - List - s.Decrypt(item.Value): %w", err)
		}
		slice[n] = item
		n++
	}

	return slice, nil
}

// Update updates a row of the data.
func (s *KeyValueService) Update(ctx context.Context, model models.KeyValue) error {
	userID := model.GetContextUserID(ctx)
	if userID == -1 {
		return fmt.Errorf("KeyValueService - Update - model.GetContextUserID(): %w", errs.ErrPermissionDenied)
	}

	if field, err := model.Validate("id", "title", "key", "value"); err != nil {
		return fmt.Errorf("KeyValueService - Update - model.Validate(): %w: %s is required", errs.ErrBadRequest, field)
	}

	model.UserID = userID
	// encrypt
	model.Key = s.cryptManager.Encrypt(model.Key)
	model.Value = s.cryptManager.Encrypt(model.Value)

	if err := s.storage.Update(ctx, model); err != nil {
		return fmt.Errorf("KeyValueService - Update - s.storage.Update(): %w", err)
	}

	return nil
}

// Delete deletes a row of the data.
func (s *KeyValueService) Delete(ctx context.Context, id int64) error {
	model := models.KeyValue{ID: id}
	userID := model.GetContextUserID(ctx)
	if userID == -1 {
		return fmt.Errorf("KeyValueService - Update - model.GetContextUserID(): %w", errs.ErrPermissionDenied)
	}
	model.UserID = userID

	// validate
	if field, err := model.Validate("user_id", "id"); err != nil {
		return fmt.Errorf("KeyValueService - Delete - model.Validate(): %w: %s is required", errs.ErrBadRequest, field)
	}

	if err := s.storage.Delete(ctx, userID, id); err != nil {
		return fmt.Errorf("KeyValueService - Delete - s.storage.Delete(): %w", err)
	}

	return nil
}

// Decrypt decrypts fields.
func (s *KeyValueService) Decrypt(field string) (string, error) {
	decrypted, err := s.cryptManager.Decrypt(field)
	if err != nil {
		return "", fmt.Errorf("KeyValueService - Decrypt - s.cryptManager.Decrypt(%s): %w", field, err)
	}

	return decrypted, nil
}
