package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/pkg/errs"
	"github.com/webkimru/go-keeper/pkg/sqlite"
)

// UserStorage contains the prepared DB connection with custom options.
type UserStorage struct {
	*sqlite.SQLite
	queryTimeout time.Duration
}

// NewUserStorage returns a new instance UserStorage with the needed options.
func NewUserStorage(sqlite *sqlite.SQLite, cfg *config.Config) *UserStorage {
	return &UserStorage{
		SQLite:       sqlite,
		queryTimeout: time.Duration(cfg.SQLite.QueryTimeout) * time.Second,
	}
}

// Add saves a new user to the database.
func (s *UserStorage) Add(ctx context.Context, user *models.User) error {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	var login string

	res := s.DB.QueryRowContext(newCtx, `
		INSERT INTO users (login, password, status) VALUES($1, $2, $3)
			ON CONFLICT (login) DO NOTHING
				RETURNING login`, user.Login, user.Password, user.Status)

	err := res.Scan(&login)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("sqlite - UserStorage - Add() - sql.ErrNoRows: %w", errs.ErrAlreadyExists)
		}

		return fmt.Errorf("sqlite - UserStorage - Add() - s.DB.QueryRowContext(): %w", err)
	}

	return nil
}

// Get looks the user for the database.
func (s *UserStorage) Get(ctx context.Context, login string) (*models.User, error) {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	var dbID int64
	var dbLogin, dbPassword string

	res := s.DB.QueryRowContext(newCtx, `
		SELECT id, login, password FROM users
			WHERE login = $1`, login)

	err := res.Scan(&dbID, &dbLogin, &dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("sqlite - UserStorage - Get() - sql.ErrNoRows: %w", errs.ErrNotFound)
		}

		return nil, fmt.Errorf("sqlite - UserStorage - Get() - QueryRowContext(): %w", err)
	}

	return &models.User{ID: dbID, Login: dbLogin, Password: dbPassword}, nil
}

// Update updates user's data in the database.
func (s *UserStorage) Update(ctx context.Context, user *models.User) error {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	row := s.DB.QueryRowContext(newCtx, `
		UPDATE users SET status = $1 WHERE login = $2`, models.UserStateRegistered, user.Login)

	err := row.Err()
	if err != nil {
		return fmt.Errorf("sqlite - UserStorage - Update() - row.Err(): %w", err)
	}

	return nil
}
