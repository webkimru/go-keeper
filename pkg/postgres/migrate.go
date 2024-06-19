package postgres

import (
	"embed"
	"fmt"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations
var migrations embed.FS

// Migrate provides up and down migrations.
func (p *Postgres) Migrate(version int64) error {
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("postgres - Migrate - goose.SetDialect(): %w", err)
	}

	db := stdlib.OpenDBFromPool(p.Pool)

	if err := goose.UpTo(db, "migrations", version); err != nil {
		return fmt.Errorf("postgres - Migrate - goose.UpTo(): %w", err)
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("postgres - Migrate - db.Close(): %w", err)
	}

	return nil
}
