package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

const (
	_defaultPingInterval   = time.Second * 1
	_defaultDataSourcePath = "data.db"
)

// SQLite contains the required options for working with the SQLite.
type SQLite struct {
	DB             *sql.DB
	PingInterval   time.Duration
	DataSourcePath string
}

// New returns the SQLite initialization.
func New(opts ...Option) (*SQLite, error) {
	s := &SQLite{
		PingInterval:   _defaultPingInterval,
		DataSourcePath: _defaultDataSourcePath,
	}
	// custom options
	for _, opt := range opts {
		opt(s)
	}

	db, err := sql.Open("sqlite", s.DataSourcePath)
	if err != nil {
		return nil, fmt.Errorf("sqlite - SQLite - Close - sql.Open(): %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.PingInterval)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("sqlite - New - db.PingContext(): %w", err)
	}

	s.DB = db

	return s, nil
}

// Close closes a DB connection.
func (s *SQLite) Close() error {
	if err := s.DB.Close(); err != nil {
		return fmt.Errorf("sqlite - SQLite - Close - s.DB.Close(): %w", err)
	}

	return nil
}
