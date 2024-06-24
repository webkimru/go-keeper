package sqlite

import (
	"embed"
	"fmt"
	"github.com/pressly/goose/v3"
)

//go:embed migrations
var migrations embed.FS

// Migrate provides up and down migrations.
func (s *SQLite) Migrate(version int64) error {
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("sqlite - Migrate - goose.SetDialect(): %w", err)
	}

	if err := goose.UpTo(s.DB, "migrations", version); err != nil {
		return fmt.Errorf("sqlite - Migrate - goose.UpTo(): %w", err)
	}

	return nil
}
