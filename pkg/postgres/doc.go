/*
Package postgres implements postgres connection.

# How to use

	pgx, err := postgres.New(cfg.PG.DatabaseDSN)

For control over its base behavior is available by creating a new option ConnectTimeout:

	pgx, err := postgres.New(
		cfg.PG.DatabaseDSN,
		postgres.ConnectTimeout(cfg.PG.ConnectTimeout),
	)

# Migrations

Package supports DB migrations. For control, create a new migration file and use the required version:

	err = pgx.Migrate(cfg.PG.MigrationVersion)
*/
package postgres
