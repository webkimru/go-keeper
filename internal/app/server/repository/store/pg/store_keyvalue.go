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

type KeyValueStorage struct {
	db           *postgres.Postgres
	queryTimeout time.Duration
}

func NewKeyValueStorage(pg *postgres.Postgres, cfg *config.Config) *KeyValueStorage {
	return &KeyValueStorage{
		db:           pg,
		queryTimeout: time.Duration(cfg.PG.QueryTimeout) * time.Second,
	}
}

// Add creates a new data.
func (s *KeyValueStorage) Add(ctx context.Context, model models.KeyValue) error {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	res, err := s.db.Pool.Query(newCtx, `
		INSERT INTO keyvalues (user_id, title, key, value)
			VALUES($1, $2, $3, $4)`, model.UserID, model.Title, model.Key, model.Value)
	if err != nil {
		return fmt.Errorf("pg - KeyValueStorage - Add() - s.db.Pool.Query(): %w", err)
	}

	// NOTICE: it's required else it will be the endless loop pool connection
	defer res.Close()

	return nil
}

// Get returns a row of the data.
func (s *KeyValueStorage) Get(ctx context.Context, id int64) (*models.KeyValue, error) {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	var dbID, dbUserID int64
	var dbTitle, dbKey, dbValue string

	res := s.db.Pool.QueryRow(newCtx, `
		SELECT id, user_id, title, key, value
			FROM keyvalues
				WHERE id = $1`, id)

	err := res.Scan(&dbID, &dbUserID, &dbTitle, &dbKey, &dbValue)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.ErrNotFound
		}

		return nil, fmt.Errorf("pg - KeyValueStorage - Get() - s.db.Pool.QueryRow(): %w", err)
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

	rows, err := s.db.Pool.Query(newCtx, `
		SELECT (id, user_id, title, key, value)
			FROM keyvalues
				WHERE user_id = $1
					LIMIT $2 OFFSET $3`, id, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("pg - KeyValueStorage - List() - s.db.Pool.Query(): %w", err)
	}

	res, err := pgx.CollectRows(rows, pgx.RowTo[models.KeyValue])
	if err != nil {
		return nil, fmt.Errorf("pg - KeyValueStorage - List() - pgx.CollectRows(): %w", err)
	}

	return res, nil
}

// Update updates a row of the data.
func (s *KeyValueStorage) Update(ctx context.Context, model models.KeyValue) error {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	tx, err := s.db.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("pg - KeyValueStorage - Update() - s.db.Pool.Begin(): %w", err)
	}
	defer tx.Rollback(ctx)

	res := tx.QueryRow(newCtx, `
		UPDATE keyvalues SET title = $1, key = $2, value = $3
			WHERE id = $4
				RETURNING user_id`, model.Title, model.Key, model.Value, model.ID)

	var dbUserID int64

	err = res.Scan(&dbUserID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errs.ErrNotFound
		}

		return fmt.Errorf("pg - KeyValueStorage - Update() - res.Scan(): %w", err)
	}

	if dbUserID != model.UserID {
		return errs.ErrPermissionDenied
	}

	return tx.Commit(ctx)
}

// Delete deletes a row of the data.
func (s *KeyValueStorage) Delete(ctx context.Context, userID, id int64) error {
	newCtx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	res, err := s.db.Pool.Exec(newCtx, `
		DELETE FROM keyvalues WHERE id = $1 and user_id = $2`, id, userID)
	if err != nil {
		return fmt.Errorf("pg - KeyValueStorage - Delete() - s.db.Pool.Exec(): %w", err)
	}

	if res.RowsAffected() == 0 {
		return errs.ErrNotFound
	}

	return nil
}
