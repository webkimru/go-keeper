package pg

import (
	"context"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/postgres"
)

type UserStorage struct {
	db *postgres.Postgres
}

func NewUserStorage(pg *postgres.Postgres) *UserStorage {
	return &UserStorage{db: pg}
}

// Add creates a new user.
func (s *UserStorage) Add(ctx context.Context, user *models.User) error {

	return nil
}

// Find look for the user.
func (s *UserStorage) Find(ctx context.Context, login string) (*models.User, error) {

	return nil, nil
}
