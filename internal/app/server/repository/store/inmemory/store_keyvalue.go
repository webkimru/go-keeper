package inmemory

import (
	"context"
	"log"
	"sync"

	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/errs"
)

// StorageKeyValue contains data in memory.
type StorageKeyValue struct {
	m        sync.RWMutex
	keyValue map[int64]map[int64]models.KeyValue
}

// NewStorageKeyValue returns a new storage for data.
func NewStorageKeyValue() *StorageKeyValue {
	return &StorageKeyValue{keyValue: make(map[int64]map[int64]models.KeyValue)}
}

// Add creates a new data.
func (s *StorageKeyValue) Add(ctx context.Context, model models.KeyValue) error {
	s.m.Lock()
	defer s.m.Unlock()

	if _, init := s.keyValue[model.UserID][1]; !init {
		s.keyValue[model.UserID] = make(map[int64]models.KeyValue)
		log.Println("init")
	}
	if _, exist := s.keyValue[model.UserID][model.ID]; exist {
		return errs.ErrAlreadyExists
	}

	idx := int64(len(s.keyValue[model.UserID]) + 1)
	model.ID = idx
	s.keyValue[model.UserID][idx] = model
	log.Printf("%v", s.keyValue)

	return nil
}

// Get returns a row of the data.
func (s *StorageKeyValue) Get(ctx context.Context, id int64) (*models.KeyValue, error) {
	userID := ctx.Value("userID").(int64)

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
func (s *StorageKeyValue) List(ctx context.Context) ([]models.KeyValue, error) {

	return nil, nil
}

// Update updates a row of the data.
func (s *StorageKeyValue) Update(ctx context.Context, model models.KeyValue) error {

	return nil
}

// Delete deletes a row of the data.
func (s *StorageKeyValue) Delete(ctx context.Context, id int64) error {

	return nil
}
