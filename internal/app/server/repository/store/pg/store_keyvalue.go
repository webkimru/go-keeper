package pg

import (
	"context"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/postgres"
)

type KeyValueStorage struct {
	db *postgres.Postgres
}

func NewKeyValueStorage(pg *postgres.Postgres) *KeyValueStorage {
	return &KeyValueStorage{db: pg}
}

// Add creates a new data.
func (s *KeyValueStorage) Add(ctx context.Context, model models.KeyValue) error {

	return nil
}

// Get returns a row of the data.
func (s *KeyValueStorage) Get(ctx context.Context, id int64) (*models.KeyValue, error) {

	return nil, nil
}

// List returns a slice of the data.
func (s *KeyValueStorage) List(ctx context.Context) ([]models.KeyValue, error) {

	return nil, nil
}

// Update updates a row of the data.
func (s *KeyValueStorage) Update(ctx context.Context, model models.KeyValue) error {

	return nil
}

// Delete deletes a row of the data.
func (s *KeyValueStorage) Delete(ctx context.Context, id int64) error {

	return nil
}
