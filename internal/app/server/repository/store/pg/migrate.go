package pg

import (
	"embed"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations
var migrations embed.FS

func Migrate(pool *pgxpool.Pool, version int64) error {
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("pg - migrate - goose.SetDialect(): %w", err)
	}

	db := stdlib.OpenDBFromPool(pool)

	if err := goose.UpTo(db, "migrations", version); err != nil {
		return fmt.Errorf("pg - migrate - goose.UpTo(): %w", err)
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("pg - migrate - db.Close(): %w", err)
	}

	return nil
}
