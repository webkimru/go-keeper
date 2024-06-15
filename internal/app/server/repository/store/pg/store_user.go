package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/webkimru/go-keeper/internal/app/server/config"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/errs"
	"github.com/webkimru/go-keeper/pkg/postgres"
)

type UserStorage struct {
	db           *postgres.Postgres
	queryTimeout time.Duration
}

func NewUserStorage(pg *postgres.Postgres, cfg *config.Config) *UserStorage {
	return &UserStorage{
		db:           pg,
		queryTimeout: time.Duration(cfg.PG.QueryTimeout) * time.Second,
	}
}

// Add creates a new user.
func (s *UserStorage) Add(ctx context.Context, user *models.User) error {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	var login string

	res := s.db.Pool.QueryRow(newCtx, `
		INSERT INTO users (login, password) VALUES($1, $2)
			ON CONFLICT (login) DO NOTHING
				RETURNING login`, user.Login, user.Password)

	err := res.Scan(&login)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errs.ErrAlreadyExists
		}

		return fmt.Errorf("pg - UserStorage - Add() - QueryRow(): %w", err)
	}

	return nil
}

// Find look for the user.
func (s *UserStorage) Find(ctx context.Context, login string) (*models.User, error) {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	var dbID int64
	var dbLogin, dbPassword string

	res := s.db.Pool.QueryRow(newCtx, `
		SELECT id, login, password FROM users
			WHERE login = $1`, login)

	err := res.Scan(&dbID, &dbLogin, &dbPassword)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.ErrNotFound
		}

		return nil, fmt.Errorf("pg - UserStorage - Find() - QueryRow(): %w", err)
	}

	return &models.User{ID: dbID, Login: dbLogin, Password: dbPassword}, nil
}
