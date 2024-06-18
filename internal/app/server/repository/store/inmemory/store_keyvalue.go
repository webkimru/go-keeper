package inmemory

import (
	"context"
	"sync"

	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/errs"
)

// KeyValueStorage contains data in memory.
type KeyValueStorage struct {
	m        sync.RWMutex
	keyValue map[int64]map[int64]models.KeyValue
}

// NewKeyValueStorage returns a new storage for data.
func NewKeyValueStorage() *KeyValueStorage {
	return &KeyValueStorage{keyValue: make(map[int64]map[int64]models.KeyValue)}
}

// Add creates a new data.
func (s *KeyValueStorage) Add(ctx context.Context, model models.KeyValue) error {
	s.m.Lock()
	defer s.m.Unlock()

	if _, init := s.keyValue[model.UserID][1]; !init {
		s.keyValue[model.UserID] = make(map[int64]models.KeyValue)
	}
	if _, exist := s.keyValue[model.UserID][model.ID]; exist {
		return errs.ErrAlreadyExists
	}

	idx := int64(len(s.keyValue[model.UserID]) + 1)
	model.ID = idx
	s.keyValue[model.UserID][idx] = model

	return nil
}

// Get returns a row of the data.
func (s *KeyValueStorage) Get(ctx context.Context, id int64) (*models.KeyValue, error) {
	userID := s.getContextUserID(ctx)
	if userID == -1 {
		return nil, errs.ErrPermissionDenied
	}

	if _, exist := s.keyValue[userID]; !exist {
		return nil, errs.ErrNotFound
	}

	data, exist := s.keyValue[userID][id]
	if !exist {
		return nil, errs.ErrNotFound
	}

	return &data, nil
}

// List returns a slice of the data.
func (s *KeyValueStorage) List(ctx context.Context, UserID, limit, offset int64) ([]models.KeyValue, error) {

	return nil, nil
}

// Update updates a row of the data.
func (s *KeyValueStorage) Update(ctx context.Context, model models.KeyValue) error {

	return nil
}

// Delete deletes a row of the data.
func (s *KeyValueStorage) Delete(ctx context.Context, UserID, id int64) error {

	return nil
}

func (s *KeyValueStorage) getContextUserID(ctx context.Context) int64 {
	id := ctx.Value(models.ContextKey("userID"))

	switch id := id.(type) {
	case int64:
		return id

	default:
		return -1
	}
}
