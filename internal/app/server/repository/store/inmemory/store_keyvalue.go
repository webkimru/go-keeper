package inmemory

import (
	"context"
	"sort"
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
	s.m.RLock()
	defer s.m.RUnlock()

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
func (s *KeyValueStorage) List(ctx context.Context, userID, limit, offset int64) ([]models.KeyValue, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	// define upper limits
	length := len(s.keyValue[userID])
	if limit > int64(length) {
		limit = int64(length)
	}
	if offset > int64(length) {
		offset = int64(length)
	}

	keys := s.sortKeyValue(userID)
	n := 0
	var slice []models.KeyValue
	for _, i := range keys {
		// take items only from offset
		if n >= int(offset) {
			slice = append(slice, s.keyValue[userID][i])
		}
		n++
		// take limited items from offset and stop
		if n-int(offset) == int(limit) {
			break
		}
	}

	return slice, nil
}

// Update updates a row of the data.
func (s *KeyValueStorage) Update(ctx context.Context, model models.KeyValue) error {
	s.m.Lock()
	defer s.m.Unlock()

	s.keyValue[model.UserID][model.ID] = model

	return nil
}

// Delete deletes a row of the data.
func (s *KeyValueStorage) Delete(ctx context.Context, userID, id int64) error {
	s.m.Lock()
	defer s.m.Unlock()

	keys := s.sortKeyValue(userID)
	slice := make(map[int64]models.KeyValue, len(keys)-1) // del: new slice len(keys)-1
	for _, i := range keys {
		if i == id {
			continue
		}
		slice[i] = s.keyValue[userID][i]
	}
	// init new slice without a deleted item
	s.keyValue[userID] = make(map[int64]models.KeyValue, len(slice))
	s.keyValue[userID] = slice

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

func (s *KeyValueStorage) sortKeyValue(userID int64) []int64 {
	length := len(s.keyValue[userID])
	keys := make([]int64, length)
	n := 0
	for k := range s.keyValue[userID] {
		keys[n] = k
		n++
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	return keys
}
