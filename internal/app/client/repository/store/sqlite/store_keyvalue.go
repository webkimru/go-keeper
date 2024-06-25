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

// KeyValueStorage contains data storage.
type KeyValueStorage struct {
	*sqlite.SQLite
	queryTimeout time.Duration
}

// NewKeyValueStorage returns a new DB connections.
func NewKeyValueStorage(sqlite *sqlite.SQLite, cfg *config.Config) *KeyValueStorage {
	return &KeyValueStorage{
		SQLite:       sqlite,
		queryTimeout: time.Duration(cfg.SQLite.QueryTimeout) * time.Second,
	}
}

// Add creates a new data.
func (s *KeyValueStorage) Add(ctx context.Context, model models.KeyValue) error {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	res, err := s.DB.QueryContext(newCtx, `
		INSERT INTO keyvalues (user_id, title, key, value, status)
			VALUES($1, $2, $3, $4, $5)`, model.UserID, model.Title, model.Key, model.Value, model.Status)
	if err != nil {
		return fmt.Errorf("sqlite - KeyValueStorage - Add() - s.DB.Query(): %w", err)
	}

	defer res.Close()

	return nil
}

// Get returns a row of the data.
func (s *KeyValueStorage) Get(ctx context.Context, id int64) (*models.KeyValue, error) {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	var dbID, dbUserID int64
	var dbTitle, dbKey, dbValue string

	res := s.DB.QueryRowContext(newCtx, `
		SELECT id, user_id, title, key, value
			FROM keyvalues
				WHERE id = $1`, id)

	err := res.Scan(&dbID, &dbUserID, &dbTitle, &dbKey, &dbValue)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("sqlite - KeyValueStorage - Get() - sql.ErrNoRows: %w", errs.ErrNotFound)
		}

		return nil, fmt.Errorf("sqlite - KeyValueStorage - Get() - s.DB.QueryRow(): %w", err)
	}

	return &models.KeyValue{
		ID:     dbID,
		UserID: dbUserID,
		Title:  dbTitle,
		Key:    dbKey,
		Value:  dbValue,
	}, nil
}

// List returns a slice of the data.
func (s *KeyValueStorage) List(ctx context.Context, id, limit, offset int64) ([]models.KeyValue, error) {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	_, err := s.DB.QueryContext(newCtx, `
		SELECT (id, user_id, title, key, value)
			FROM keyvalues
				WHERE user_id = $1
					LIMIT $2 OFFSET $3`, id, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("sqlite - KeyValueStorage - List() - s.DB.Query(): %w", err)
	}

	// TODO: list

	return nil, nil
}

// Update updates a row of the data.
func (s *KeyValueStorage) Update(ctx context.Context, model models.KeyValue) error {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	tx, err := s.DB.BeginTx(newCtx, nil)
	if err != nil {
		return fmt.Errorf("sqlite - KeyValueStorage - Update() - s.DB.Begin(): %w", err)
	}
	defer tx.Rollback()

	res := tx.QueryRow(`
		UPDATE keyvalues SET title = $1, key = $2, value = $3
			WHERE id = $4
				RETURNING user_id`, model.Title, model.Key, model.Value, model.ID)

	var dbUserID int64

	err = res.Scan(&dbUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("sqlite - KeyValueStorage - Update() - sql.ErrNoRows: %w", errs.ErrNotFound)
		}

		return fmt.Errorf("sqlite - KeyValueStorage - Update() - res.Scan(): %w", err)
	}

	if dbUserID != model.UserID {
		return fmt.Errorf("sqlite - KeyValueStorage - Update() - dbUserID is not equal model.UserID: %w", errs.ErrPermissionDenied)
	}

	return tx.Commit()
}

// Delete deletes a row of the data.
func (s *KeyValueStorage) Delete(ctx context.Context, userID, id int64) error {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	res, err := s.DB.ExecContext(newCtx, `
		DELETE FROM keyvalues WHERE id = $1 and user_id = $2`, id, userID)
	if err != nil {
		return fmt.Errorf("sqlite - KeyValueStorage - Delete() - s.DB.Exec(): %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("sqlite - KeyValueStorage - Delete() - res.RowsAffected(): %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("sqlite - KeyValueStorage - Delete() - res.RowsAffected(): %w", errs.ErrNotFound)
	}

	return nil
}
