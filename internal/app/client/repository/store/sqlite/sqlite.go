package sqlite

import (
	"context"
	"time"

	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/pkg/sqlite"
)

// UserStorage contains the prepared DB connection with custom options.
type UserStorage struct {
	db           *sqlite.SQLite
	pingInterval time.Duration
}

// NewUserStorage returns a new instance UserStorage with the needed options.
func NewUserStorage(sqlite *sqlite.SQLite, cfg *config.Config) *UserStorage {
	return &UserStorage{
		db:           sqlite,
		pingInterval: time.Duration(cfg.SQLite.PingInterval) * time.Second,
	}
}

// Add saves a new user to the database.
func (s *UserStorage) Add(ctx context.Context, user *models.User) error {

	return nil
}

// Get looks the user for the database.
func (s *UserStorage) Get(ctx context.Context, login string) (*models.User, error) {

	return nil, nil
}
